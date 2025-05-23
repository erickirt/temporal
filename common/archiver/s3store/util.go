package s3store

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	commonpb "go.temporal.io/api/common/v1"
	historypb "go.temporal.io/api/history/v1"
	"go.temporal.io/api/serviceerror"
	workflowpb "go.temporal.io/api/workflow/v1"
	archiverspb "go.temporal.io/server/api/archiver/v1"
	"go.temporal.io/server/common/archiver"
	"go.temporal.io/server/common/codec"
	"go.temporal.io/server/common/searchattribute"
	"go.uber.org/multierr"
	"google.golang.org/protobuf/proto"
)

// encoding & decoding util

func Encode(message proto.Message) ([]byte, error) {
	encoder := codec.NewJSONPBEncoder()
	return encoder.Encode(message)
}

func decodeVisibilityRecord(data []byte) (*archiverspb.VisibilityRecord, error) {
	record := &archiverspb.VisibilityRecord{}
	encoder := codec.NewJSONPBEncoder()
	err := encoder.Decode(data, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func SerializeToken(token interface{}) ([]byte, error) {
	if token == nil {
		return nil, nil
	}
	return json.Marshal(token)
}

func deserializeGetHistoryToken(bytes []byte) (*getHistoryToken, error) {
	token := &getHistoryToken{}
	err := json.Unmarshal(bytes, token)
	return token, err
}

func deserializeQueryVisibilityToken(bytes []byte) *string {
	var ret = string(bytes)
	return &ret
}
func serializeQueryVisibilityToken(token string) []byte {
	return []byte(token)
}

// Only validates the scheme and buckets are passed
func SoftValidateURI(URI archiver.URI) error {
	if URI.Scheme() != URIScheme {
		return archiver.ErrURISchemeMismatch
	}
	if len(URI.Hostname()) == 0 {
		return errNoBucketSpecified
	}
	return nil
}

func BucketExists(ctx context.Context, s3cli s3iface.S3API, URI archiver.URI) error {
	ctx, cancel := ensureContextTimeout(ctx)
	defer cancel()
	_, err := s3cli.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(URI.Hostname()),
	})
	if err == nil {
		return nil
	}
	if IsNotFoundError(err) {
		return errBucketNotExists
	}
	return err
}

func KeyExists(ctx context.Context, s3cli s3iface.S3API, URI archiver.URI, key string) (bool, error) {
	ctx, cancel := ensureContextTimeout(ctx)
	defer cancel()
	_, err := s3cli.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(URI.Hostname()),
		Key:    aws.String(key),
	})
	if err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsNotFoundError(err error) bool {
	aerr, ok := err.(awserr.Error)
	return ok && (aerr.Code() == "NotFound")
}

// Key construction
func constructHistoryKey(path, namespaceID, workflowID, runID string, version int64, batchIdx int) string {
	prefix := constructHistoryKeyPrefixWithVersion(path, namespaceID, workflowID, runID, version)
	return fmt.Sprintf("%s%d", prefix, batchIdx)
}

func constructHistoryKeyPrefixWithVersion(path, namespaceID, workflowID, runID string, version int64) string {
	prefix := constructHistoryKeyPrefix(path, namespaceID, workflowID, runID)
	return fmt.Sprintf("%s/%v/", prefix, version)
}

func constructHistoryKeyPrefix(path, namespaceID, workflowID, runID string) string {
	return strings.TrimLeft(strings.Join([]string{path, namespaceID, "history", workflowID, runID}, "/"), "/")
}

func constructTimeBasedSearchKey(path, namespaceID, primaryIndexKey, primaryIndexValue, secondaryIndexKey string, t time.Time, precision string) string {
	var timeFormat = ""
	switch precision {
	case PrecisionSecond:
		timeFormat = ":05"
		fallthrough
	case PrecisionMinute:
		timeFormat = ":04" + timeFormat
		fallthrough
	case PrecisionHour:
		timeFormat = "15" + timeFormat
		fallthrough
	case PrecisionDay:
		timeFormat = "2006-01-02T" + timeFormat
	}

	return fmt.Sprintf(
		"%s/%s",
		constructIndexedVisibilitySearchPrefix(path, namespaceID, primaryIndexKey, primaryIndexValue, secondaryIndexKey),
		t.Format(timeFormat),
	)
}

func constructTimestampIndex(path, namespaceID, primaryIndexKey, primaryIndexValue, secondaryIndexKey string, secondaryIndexValue time.Time, runID string) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		constructIndexedVisibilitySearchPrefix(path, namespaceID, primaryIndexKey, primaryIndexValue, secondaryIndexKey),
		secondaryIndexValue.Format(time.RFC3339),
		runID,
	)
}

func constructIndexedVisibilitySearchPrefix(
	path string,
	namespaceID string,
	primaryIndexKey string,
	primaryIndexValue string,
	secondaryIndexType string,
) string {
	return strings.TrimLeft(
		strings.Join(
			[]string{path, namespaceID, "visibility", primaryIndexKey, primaryIndexValue, secondaryIndexType},
			"/",
		),
		"/",
	)
}

func constructVisibilitySearchPrefix(path, namespaceID string) string {
	return strings.TrimLeft(strings.Join([]string{path, namespaceID, "visibility"}, "/"), "/")
}

func ensureContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if _, ok := ctx.Deadline(); ok {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, defaultBlobstoreTimeout)
}
func Upload(ctx context.Context, s3cli s3iface.S3API, URI archiver.URI, key string, data []byte) error {
	ctx, cancel := ensureContextTimeout(ctx)
	defer cancel()

	_, err := s3cli.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(URI.Hostname()),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchBucket {
				return serviceerror.NewInvalidArgument(errBucketNotExists.Error())
			}
		}
		return err
	}
	return nil
}

func Download(ctx context.Context, s3cli s3iface.S3API, URI archiver.URI, key string) ([]byte, error) {
	ctx, cancel := ensureContextTimeout(ctx)
	defer cancel()
	result, err := s3cli.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(URI.Hostname()),
		Key:    aws.String(key),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchBucket {
				return nil, serviceerror.NewInvalidArgument(errBucketNotExists.Error())
			}

			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return nil, serviceerror.NewNotFound(archiver.ErrHistoryNotExist.Error())
			}
		}
		return nil, err
	}

	defer func() {
		if ierr := result.Body.Close(); ierr != nil {
			err = multierr.Append(err, ierr)
		}
	}()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func historyMutated(request *archiver.ArchiveHistoryRequest, historyBatches []*historypb.History, isLast bool) bool {
	lastBatch := historyBatches[len(historyBatches)-1].Events
	lastEvent := lastBatch[len(lastBatch)-1]
	lastFailoverVersion := lastEvent.GetVersion()
	if lastFailoverVersion > request.CloseFailoverVersion {
		return true
	}

	if !isLast {
		return false
	}
	lastEventID := lastEvent.GetEventId()
	return lastFailoverVersion != request.CloseFailoverVersion || lastEventID+1 != request.NextEventID
}

func convertToExecutionInfo(record *archiverspb.VisibilityRecord, saTypeMap searchattribute.NameTypeMap) (*workflowpb.WorkflowExecutionInfo, error) {
	searchAttributes, err := searchattribute.Parse(record.SearchAttributes, &saTypeMap)
	if err != nil {
		return nil, err
	}

	return &workflowpb.WorkflowExecutionInfo{
		Execution: &commonpb.WorkflowExecution{
			WorkflowId: record.GetWorkflowId(),
			RunId:      record.GetRunId(),
		},
		Type: &commonpb.WorkflowType{
			Name: record.WorkflowTypeName,
		},
		StartTime:         record.StartTime,
		ExecutionTime:     record.ExecutionTime,
		CloseTime:         record.CloseTime,
		ExecutionDuration: record.ExecutionDuration,
		Status:            record.Status,
		HistoryLength:     record.HistoryLength,
		Memo:              record.Memo,
		SearchAttributes:  searchAttributes,
	}, nil
}
