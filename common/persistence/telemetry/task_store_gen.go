// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by gowrap. DO NOT EDIT.
// template: gowrap_template
// gowrap: http://github.com/hexdigest/gowrap

package telemetry

//go:generate gowrap gen -p go.temporal.io/server/common/persistence -i TaskStore -t gowrap_template -o task_store_gen.go -l ""

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	_sourcePersistence "go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/telemetry"
)

// telemetryTaskStore implements TaskStore interface instrumented with OpenTelemetry.
type telemetryTaskStore struct {
	_sourcePersistence.TaskStore
	tracer    trace.Tracer
	logger    log.Logger
	debugMode bool
}

// newTelemetryTaskStore returns telemetryTaskStore.
func newTelemetryTaskStore(
	base _sourcePersistence.TaskStore,
	logger log.Logger,
	tracer trace.Tracer,
) telemetryTaskStore {
	return telemetryTaskStore{
		TaskStore: base,
		tracer:    tracer,
		debugMode: telemetry.DebugMode(),
	}
}

// CompleteTasksLessThan wraps TaskStore.CompleteTasksLessThan.
func (d telemetryTaskStore) CompleteTasksLessThan(ctx context.Context, request *_sourcePersistence.CompleteTasksLessThanRequest) (i1 int, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/CompleteTasksLessThan")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("CompleteTasksLessThan"))

	i1, err = d.TaskStore.CompleteTasksLessThan(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.CompleteTasksLessThanRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(i1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize int for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// CountTaskQueuesByBuildId wraps TaskStore.CountTaskQueuesByBuildId.
func (d telemetryTaskStore) CountTaskQueuesByBuildId(ctx context.Context, request *_sourcePersistence.CountTaskQueuesByBuildIdRequest) (i1 int, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/CountTaskQueuesByBuildId")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("CountTaskQueuesByBuildId"))

	i1, err = d.TaskStore.CountTaskQueuesByBuildId(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.CountTaskQueuesByBuildIdRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(i1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize int for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// CreateTaskQueue wraps TaskStore.CreateTaskQueue.
func (d telemetryTaskStore) CreateTaskQueue(ctx context.Context, request *_sourcePersistence.InternalCreateTaskQueueRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/CreateTaskQueue")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("CreateTaskQueue"))

	err = d.TaskStore.CreateTaskQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalCreateTaskQueueRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

	}

	return
}

// CreateTasks wraps TaskStore.CreateTasks.
func (d telemetryTaskStore) CreateTasks(ctx context.Context, request *_sourcePersistence.InternalCreateTasksRequest) (cp1 *_sourcePersistence.CreateTasksResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/CreateTasks")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("CreateTasks"))

	cp1, err = d.TaskStore.CreateTasks(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalCreateTasksRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(cp1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.CreateTasksResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// DeleteTaskQueue wraps TaskStore.DeleteTaskQueue.
func (d telemetryTaskStore) DeleteTaskQueue(ctx context.Context, request *_sourcePersistence.DeleteTaskQueueRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/DeleteTaskQueue")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("DeleteTaskQueue"))

	err = d.TaskStore.DeleteTaskQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.DeleteTaskQueueRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

	}

	return
}

// GetTaskQueue wraps TaskStore.GetTaskQueue.
func (d telemetryTaskStore) GetTaskQueue(ctx context.Context, request *_sourcePersistence.InternalGetTaskQueueRequest) (ip1 *_sourcePersistence.InternalGetTaskQueueResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/GetTaskQueue")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("GetTaskQueue"))

	ip1, err = d.TaskStore.GetTaskQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalGetTaskQueueRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalGetTaskQueueResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// GetTaskQueueUserData wraps TaskStore.GetTaskQueueUserData.
func (d telemetryTaskStore) GetTaskQueueUserData(ctx context.Context, request *_sourcePersistence.GetTaskQueueUserDataRequest) (ip1 *_sourcePersistence.InternalGetTaskQueueUserDataResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/GetTaskQueueUserData")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("GetTaskQueueUserData"))

	ip1, err = d.TaskStore.GetTaskQueueUserData(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.GetTaskQueueUserDataRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalGetTaskQueueUserDataResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// GetTaskQueuesByBuildId wraps TaskStore.GetTaskQueuesByBuildId.
func (d telemetryTaskStore) GetTaskQueuesByBuildId(ctx context.Context, request *_sourcePersistence.GetTaskQueuesByBuildIdRequest) (sa1 []string, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/GetTaskQueuesByBuildId")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("GetTaskQueuesByBuildId"))

	sa1, err = d.TaskStore.GetTaskQueuesByBuildId(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.GetTaskQueuesByBuildIdRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(sa1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize []string for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// GetTasks wraps TaskStore.GetTasks.
func (d telemetryTaskStore) GetTasks(ctx context.Context, request *_sourcePersistence.GetTasksRequest) (ip1 *_sourcePersistence.InternalGetTasksResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/GetTasks")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("GetTasks"))

	ip1, err = d.TaskStore.GetTasks(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.GetTasksRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalGetTasksResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// ListTaskQueue wraps TaskStore.ListTaskQueue.
func (d telemetryTaskStore) ListTaskQueue(ctx context.Context, request *_sourcePersistence.ListTaskQueueRequest) (ip1 *_sourcePersistence.InternalListTaskQueueResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/ListTaskQueue")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("ListTaskQueue"))

	ip1, err = d.TaskStore.ListTaskQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.ListTaskQueueRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalListTaskQueueResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// ListTaskQueueUserDataEntries wraps TaskStore.ListTaskQueueUserDataEntries.
func (d telemetryTaskStore) ListTaskQueueUserDataEntries(ctx context.Context, request *_sourcePersistence.ListTaskQueueUserDataEntriesRequest) (ip1 *_sourcePersistence.InternalListTaskQueueUserDataEntriesResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/ListTaskQueueUserDataEntries")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("ListTaskQueueUserDataEntries"))

	ip1, err = d.TaskStore.ListTaskQueueUserDataEntries(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.ListTaskQueueUserDataEntriesRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalListTaskQueueUserDataEntriesResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// UpdateTaskQueue wraps TaskStore.UpdateTaskQueue.
func (d telemetryTaskStore) UpdateTaskQueue(ctx context.Context, request *_sourcePersistence.InternalUpdateTaskQueueRequest) (up1 *_sourcePersistence.UpdateTaskQueueResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/UpdateTaskQueue")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("UpdateTaskQueue"))

	up1, err = d.TaskStore.UpdateTaskQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalUpdateTaskQueueRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(up1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.UpdateTaskQueueResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// UpdateTaskQueueUserData wraps TaskStore.UpdateTaskQueueUserData.
func (d telemetryTaskStore) UpdateTaskQueueUserData(ctx context.Context, request *_sourcePersistence.InternalUpdateTaskQueueUserDataRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.TaskStore/UpdateTaskQueueUserData")
	defer span.End()

	span.SetAttributes(attribute.Key("persistence.store").String("TaskStore"))
	span.SetAttributes(attribute.Key("persistence.method").String("UpdateTaskQueueUserData"))

	err = d.TaskStore.UpdateTaskQueueUserData(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalUpdateTaskQueueUserDataRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

	}

	return
}
