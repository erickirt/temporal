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

//go:generate gowrap gen -p go.temporal.io/server/common/persistence -i NexusEndpointStore -t gowrap_template -o nexus_endpoint_store_gen.go -l ""

import (
	"context"
	"encoding/json"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	_sourcePersistence "go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/telemetry"
)

// telemetryNexusEndpointStore implements NexusEndpointStore interface instrumented with OpenTelemetry.
type telemetryNexusEndpointStore struct {
	_sourcePersistence.NexusEndpointStore
	tracer    trace.Tracer
	logger    log.Logger
	debugMode bool
}

// newTelemetryNexusEndpointStore returns telemetryNexusEndpointStore.
func newTelemetryNexusEndpointStore(
	base _sourcePersistence.NexusEndpointStore,
	logger log.Logger,
	tracer trace.Tracer,
) telemetryNexusEndpointStore {
	return telemetryNexusEndpointStore{
		NexusEndpointStore: base,
		tracer:             tracer,
		debugMode:          telemetry.DebugMode(),
	}
}

// CreateOrUpdateNexusEndpoint wraps NexusEndpointStore.CreateOrUpdateNexusEndpoint.
func (d telemetryNexusEndpointStore) CreateOrUpdateNexusEndpoint(ctx context.Context, request *_sourcePersistence.InternalCreateOrUpdateNexusEndpointRequest) (err error) {
	ctx, span := d.tracer.Start(
		ctx,
		"persistence.NexusEndpointStore/CreateOrUpdateNexusEndpoint",
		trace.WithAttributes(
			attribute.Key("persistence.store").String("NexusEndpointStore"),
			attribute.Key("persistence.method").String("CreateOrUpdateNexusEndpoint"),
		))
	defer span.End()

	if deadline, ok := ctx.Deadline(); ok {
		span.SetAttributes(attribute.String("deadline", deadline.Format(time.RFC3339Nano)))
		span.SetAttributes(attribute.String("timeout", time.Until(deadline).String()))
	}

	err = d.NexusEndpointStore.CreateOrUpdateNexusEndpoint(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalCreateOrUpdateNexusEndpointRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

	}

	return
}

// DeleteNexusEndpoint wraps NexusEndpointStore.DeleteNexusEndpoint.
func (d telemetryNexusEndpointStore) DeleteNexusEndpoint(ctx context.Context, request *_sourcePersistence.DeleteNexusEndpointRequest) (err error) {
	ctx, span := d.tracer.Start(
		ctx,
		"persistence.NexusEndpointStore/DeleteNexusEndpoint",
		trace.WithAttributes(
			attribute.Key("persistence.store").String("NexusEndpointStore"),
			attribute.Key("persistence.method").String("DeleteNexusEndpoint"),
		))
	defer span.End()

	if deadline, ok := ctx.Deadline(); ok {
		span.SetAttributes(attribute.String("deadline", deadline.Format(time.RFC3339Nano)))
		span.SetAttributes(attribute.String("timeout", time.Until(deadline).String()))
	}

	err = d.NexusEndpointStore.DeleteNexusEndpoint(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.DeleteNexusEndpointRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

	}

	return
}

// GetNexusEndpoint wraps NexusEndpointStore.GetNexusEndpoint.
func (d telemetryNexusEndpointStore) GetNexusEndpoint(ctx context.Context, request *_sourcePersistence.GetNexusEndpointRequest) (ip1 *_sourcePersistence.InternalNexusEndpoint, err error) {
	ctx, span := d.tracer.Start(
		ctx,
		"persistence.NexusEndpointStore/GetNexusEndpoint",
		trace.WithAttributes(
			attribute.Key("persistence.store").String("NexusEndpointStore"),
			attribute.Key("persistence.method").String("GetNexusEndpoint"),
		))
	defer span.End()

	if deadline, ok := ctx.Deadline(); ok {
		span.SetAttributes(attribute.String("deadline", deadline.Format(time.RFC3339Nano)))
		span.SetAttributes(attribute.String("timeout", time.Until(deadline).String()))
	}

	ip1, err = d.NexusEndpointStore.GetNexusEndpoint(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.GetNexusEndpointRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalNexusEndpoint for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}

// ListNexusEndpoints wraps NexusEndpointStore.ListNexusEndpoints.
func (d telemetryNexusEndpointStore) ListNexusEndpoints(ctx context.Context, request *_sourcePersistence.ListNexusEndpointsRequest) (ip1 *_sourcePersistence.InternalListNexusEndpointsResponse, err error) {
	ctx, span := d.tracer.Start(
		ctx,
		"persistence.NexusEndpointStore/ListNexusEndpoints",
		trace.WithAttributes(
			attribute.Key("persistence.store").String("NexusEndpointStore"),
			attribute.Key("persistence.method").String("ListNexusEndpoints"),
		))
	defer span.End()

	if deadline, ok := ctx.Deadline(); ok {
		span.SetAttributes(attribute.String("deadline", deadline.Format(time.RFC3339Nano)))
		span.SetAttributes(attribute.String("timeout", time.Until(deadline).String()))
	}

	ip1, err = d.NexusEndpointStore.ListNexusEndpoints(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	if d.debugMode {

		requestPayload, err := json.MarshalIndent(request, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.ListNexusEndpointsRequest for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.request.payload").String(string(requestPayload)))
		}

		responsePayload, err := json.MarshalIndent(ip1, "", "    ")
		if err != nil {
			d.logger.Error("failed to serialize *_sourcePersistence.InternalListNexusEndpointsResponse for OTEL span", tag.Error(err))
		} else {
			span.SetAttributes(attribute.Key("persistence.response.payload").String(string(responsePayload)))
		}

	}

	return
}
