package replication

import (
	"context"
	"sync"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/api/adminservice/v1"
	enumsspb "go.temporal.io/server/api/enums/v1"
	replicationspb "go.temporal.io/server/api/replication/v1"
	"go.temporal.io/server/client"
	"go.temporal.io/server/common/headers"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/namespace/nsreplication"
	"go.temporal.io/server/common/persistence"
)

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination eager_namespace_refresher_mock.go

type (
	EagerNamespaceRefresher interface {
		UpdateNamespaceFailoverVersion(namespaceId namespace.ID, targetFailoverVersion int64) error
		SyncNamespaceFromSourceCluster(ctx context.Context, namespaceId namespace.ID, sourceCluster string) (*namespace.Namespace, error)
	}

	eagerNamespaceRefresherImpl struct {
		metadataManager         persistence.MetadataManager
		namespaceRegistry       namespace.Registry
		logger                  log.Logger
		lock                    sync.Mutex
		clientBean              client.Bean
		replicationTaskExecutor nsreplication.TaskExecutor
		currentCluster          string
		metricsHandler          metrics.Handler
	}
)

func NewEagerNamespaceRefresher(
	metadataManager persistence.MetadataManager,
	namespaceRegistry namespace.Registry,
	logger log.Logger,
	clientBean client.Bean,
	replicationTaskExecutor nsreplication.TaskExecutor,
	currentCluster string,
	metricsHandler metrics.Handler) EagerNamespaceRefresher {
	return &eagerNamespaceRefresherImpl{
		metadataManager:         metadataManager,
		namespaceRegistry:       namespaceRegistry,
		logger:                  logger,
		clientBean:              clientBean,
		replicationTaskExecutor: replicationTaskExecutor,
		currentCluster:          currentCluster,
		metricsHandler:          metricsHandler,
	}
}

func (e *eagerNamespaceRefresherImpl) UpdateNamespaceFailoverVersion(namespaceId namespace.ID, targetFailoverVersion int64) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	ns, err := e.namespaceRegistry.GetNamespaceByID(namespaceId)
	switch err.(type) {
	case nil:
	case *serviceerror.NamespaceNotFound:
		// TODO: Handle NamespaceNotFound case, probably retrieve the namespace from the source cluster?
		return nil
	default:
		// do nothing as this is the best effort to update the namespace
		e.logger.Debug("Failed to get namespace from registry", tag.Error(err))
		return err
	}

	if ns.FailoverVersion() >= targetFailoverVersion {
		return nil
	}

	ctx := headers.SetCallerInfo(context.TODO(), headers.SystemPreemptableCallerInfo)
	resp, err := e.metadataManager.GetNamespace(ctx, &persistence.GetNamespaceRequest{
		ID: namespaceId.String(),
	})
	if err != nil {
		e.logger.Debug("Failed to get namespace from persistent", tag.Error(err))
		return err
	}

	currentFailoverVersion := resp.Namespace.FailoverVersion
	if currentFailoverVersion >= targetFailoverVersion {
		// DB may have a fresher version of namespace, so compare again
		return nil
	}

	metadata, err := e.metadataManager.GetMetadata(ctx)
	if err != nil {
		e.logger.Debug("Failed to get metadata", tag.Error(err))
		return err
	}

	request := &persistence.UpdateNamespaceRequest{
		Namespace:           resp.Namespace,
		NotificationVersion: metadata.NotificationVersion,
		IsGlobalNamespace:   resp.IsGlobalNamespace,
	}

	request.Namespace.FailoverVersion = targetFailoverVersion
	request.Namespace.FailoverNotificationVersion = metadata.NotificationVersion

	// Question: is it ok to only update failover version WITHOUT updating FailoverHistory?
	// request.Namespace.ReplicationConfig.FailoverHistory = ??

	if err := e.metadataManager.UpdateNamespace(ctx, request); err != nil {
		e.logger.Info("Failed to update namespace", tag.Error(err))
		return err
	}
	return nil
}

func (e *eagerNamespaceRefresherImpl) SyncNamespaceFromSourceCluster(
	ctx context.Context,
	namespaceId namespace.ID,
	sourceCluster string,
) (*namespace.Namespace, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	adminClient, err := e.clientBean.GetRemoteAdminClient(sourceCluster)
	if err != nil {
		return nil, err
	}
	resp, err := adminClient.GetNamespace(ctx, &adminservice.GetNamespaceRequest{
		Attributes: &adminservice.GetNamespaceRequest_Id{
			Id: namespaceId.String(),
		},
	})
	if err != nil {
		return nil, err
	}
	if !resp.GetIsGlobalNamespace() {
		return nil, serviceerror.NewFailedPreconditionf("Not a global namespace: %v", namespaceId)
	}
	hasCurrentCluster := false
	for _, c := range resp.GetReplicationConfig().GetClusters() {
		if e.currentCluster == c.GetClusterName() {
			hasCurrentCluster = true
		}
	}
	if !hasCurrentCluster {
		metrics.ReplicationOutlierNamespace.With(e.metricsHandler).Record(1)
		return nil, serviceerror.NewFailedPrecondition("Namespace does not belong to current cluster")
	}
	_, err = e.namespaceRegistry.GetNamespaceByID(namespaceId)
	var operation enumsspb.NamespaceOperation
	switch err.(type) {
	case *serviceerror.NamespaceNotFound:
		operation = enumsspb.NAMESPACE_OPERATION_CREATE
	case nil:
		operation = enumsspb.NAMESPACE_OPERATION_UPDATE
	default:
		return nil, err
	}
	task := &replicationspb.NamespaceTaskAttributes{
		NamespaceOperation: operation,
		Id:                 resp.GetInfo().Id,
		Info:               resp.GetInfo(),
		Config:             resp.GetConfig(),
		ReplicationConfig:  resp.GetReplicationConfig(),
		ConfigVersion:      resp.GetConfigVersion(),
		FailoverVersion:    resp.GetFailoverVersion(),
		FailoverHistory:    resp.GetFailoverHistory(),
	}
	err = e.replicationTaskExecutor.Execute(ctx, task)
	if err != nil {
		return nil, err
	}
	return e.namespaceRegistry.RefreshNamespaceById(namespaceId)
}
