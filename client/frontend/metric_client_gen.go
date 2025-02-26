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

// Code generated by cmd/tools/rpcwrappers. DO NOT EDIT.

package frontend

import (
	"context"

	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"

	"go.temporal.io/server/common/metrics"
)

func (c *metricClient) CountWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.CountWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.CountWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientCountWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.CountWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) CreateSchedule(
	ctx context.Context,
	request *workflowservice.CreateScheduleRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.CreateScheduleResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientCreateScheduleScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.CreateSchedule(ctx, request, opts...)
}

func (c *metricClient) DeleteSchedule(
	ctx context.Context,
	request *workflowservice.DeleteScheduleRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DeleteScheduleResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDeleteScheduleScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DeleteSchedule(ctx, request, opts...)
}

func (c *metricClient) DeprecateNamespace(
	ctx context.Context,
	request *workflowservice.DeprecateNamespaceRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DeprecateNamespaceResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDeprecateNamespaceScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DeprecateNamespace(ctx, request, opts...)
}

func (c *metricClient) DescribeBatchOperation(
	ctx context.Context,
	request *workflowservice.DescribeBatchOperationRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DescribeBatchOperationResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDescribeBatchOperationScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DescribeBatchOperation(ctx, request, opts...)
}

func (c *metricClient) DescribeNamespace(
	ctx context.Context,
	request *workflowservice.DescribeNamespaceRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DescribeNamespaceResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDescribeNamespaceScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DescribeNamespace(ctx, request, opts...)
}

func (c *metricClient) DescribeSchedule(
	ctx context.Context,
	request *workflowservice.DescribeScheduleRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DescribeScheduleResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDescribeScheduleScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DescribeSchedule(ctx, request, opts...)
}

func (c *metricClient) DescribeTaskQueue(
	ctx context.Context,
	request *workflowservice.DescribeTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DescribeTaskQueueResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDescribeTaskQueueScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DescribeTaskQueue(ctx, request, opts...)
}

func (c *metricClient) DescribeWorkflowExecution(
	ctx context.Context,
	request *workflowservice.DescribeWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.DescribeWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientDescribeWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.DescribeWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) GetClusterInfo(
	ctx context.Context,
	request *workflowservice.GetClusterInfoRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetClusterInfoResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetClusterInfoScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetClusterInfo(ctx, request, opts...)
}

func (c *metricClient) GetSearchAttributes(
	ctx context.Context,
	request *workflowservice.GetSearchAttributesRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetSearchAttributesResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetSearchAttributesScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetSearchAttributes(ctx, request, opts...)
}

func (c *metricClient) GetSystemInfo(
	ctx context.Context,
	request *workflowservice.GetSystemInfoRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetSystemInfoResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetSystemInfoScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetSystemInfo(ctx, request, opts...)
}

func (c *metricClient) GetWorkerBuildIdOrdering(
	ctx context.Context,
	request *workflowservice.GetWorkerBuildIdOrderingRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetWorkerBuildIdOrderingResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetWorkerBuildIdOrderingScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetWorkerBuildIdOrdering(ctx, request, opts...)
}

func (c *metricClient) GetWorkflowExecutionHistory(
	ctx context.Context,
	request *workflowservice.GetWorkflowExecutionHistoryRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetWorkflowExecutionHistoryResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetWorkflowExecutionHistoryScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetWorkflowExecutionHistory(ctx, request, opts...)
}

func (c *metricClient) GetWorkflowExecutionHistoryReverse(
	ctx context.Context,
	request *workflowservice.GetWorkflowExecutionHistoryReverseRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.GetWorkflowExecutionHistoryReverseResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientGetWorkflowExecutionHistoryReverseScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.GetWorkflowExecutionHistoryReverse(ctx, request, opts...)
}

func (c *metricClient) ListArchivedWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.ListArchivedWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListArchivedWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListArchivedWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListArchivedWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) ListBatchOperations(
	ctx context.Context,
	request *workflowservice.ListBatchOperationsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListBatchOperationsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListBatchOperationsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListBatchOperations(ctx, request, opts...)
}

func (c *metricClient) ListClosedWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.ListClosedWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListClosedWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListClosedWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListClosedWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) ListNamespaces(
	ctx context.Context,
	request *workflowservice.ListNamespacesRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListNamespacesResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListNamespacesScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListNamespaces(ctx, request, opts...)
}

func (c *metricClient) ListOpenWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.ListOpenWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListOpenWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListOpenWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListOpenWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) ListScheduleMatchingTimes(
	ctx context.Context,
	request *workflowservice.ListScheduleMatchingTimesRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListScheduleMatchingTimesResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListScheduleMatchingTimesScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListScheduleMatchingTimes(ctx, request, opts...)
}

func (c *metricClient) ListSchedules(
	ctx context.Context,
	request *workflowservice.ListSchedulesRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListSchedulesResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListSchedulesScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListSchedules(ctx, request, opts...)
}

func (c *metricClient) ListTaskQueuePartitions(
	ctx context.Context,
	request *workflowservice.ListTaskQueuePartitionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListTaskQueuePartitionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListTaskQueuePartitionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListTaskQueuePartitions(ctx, request, opts...)
}

func (c *metricClient) ListWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.ListWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ListWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientListWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ListWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) PatchSchedule(
	ctx context.Context,
	request *workflowservice.PatchScheduleRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.PatchScheduleResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientPatchScheduleScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.PatchSchedule(ctx, request, opts...)
}

func (c *metricClient) PollActivityTaskQueue(
	ctx context.Context,
	request *workflowservice.PollActivityTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.PollActivityTaskQueueResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientPollActivityTaskQueueScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.PollActivityTaskQueue(ctx, request, opts...)
}

func (c *metricClient) PollWorkflowTaskQueue(
	ctx context.Context,
	request *workflowservice.PollWorkflowTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.PollWorkflowTaskQueueResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientPollWorkflowTaskQueueScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.PollWorkflowTaskQueue(ctx, request, opts...)
}

func (c *metricClient) QueryWorkflow(
	ctx context.Context,
	request *workflowservice.QueryWorkflowRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.QueryWorkflowResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientQueryWorkflowScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.QueryWorkflow(ctx, request, opts...)
}

func (c *metricClient) RecordActivityTaskHeartbeat(
	ctx context.Context,
	request *workflowservice.RecordActivityTaskHeartbeatRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RecordActivityTaskHeartbeatResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRecordActivityTaskHeartbeatScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RecordActivityTaskHeartbeat(ctx, request, opts...)
}

func (c *metricClient) RecordActivityTaskHeartbeatById(
	ctx context.Context,
	request *workflowservice.RecordActivityTaskHeartbeatByIdRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RecordActivityTaskHeartbeatByIdResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRecordActivityTaskHeartbeatByIdScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RecordActivityTaskHeartbeatById(ctx, request, opts...)
}

func (c *metricClient) RegisterNamespace(
	ctx context.Context,
	request *workflowservice.RegisterNamespaceRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RegisterNamespaceResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRegisterNamespaceScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RegisterNamespace(ctx, request, opts...)
}

func (c *metricClient) RequestCancelWorkflowExecution(
	ctx context.Context,
	request *workflowservice.RequestCancelWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RequestCancelWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRequestCancelWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RequestCancelWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) ResetStickyTaskQueue(
	ctx context.Context,
	request *workflowservice.ResetStickyTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ResetStickyTaskQueueResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientResetStickyTaskQueueScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ResetStickyTaskQueue(ctx, request, opts...)
}

func (c *metricClient) ResetWorkflowExecution(
	ctx context.Context,
	request *workflowservice.ResetWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ResetWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientResetWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ResetWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskCanceled(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskCanceledRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskCanceledResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskCanceledScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskCanceled(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskCanceledById(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskCanceledByIdRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskCanceledByIdResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskCanceledByIdScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskCanceledById(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskCompleted(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskCompletedResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskCompletedScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskCompletedById(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskCompletedByIdRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskCompletedByIdResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskCompletedByIdScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskCompletedById(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskFailed(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskFailedRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskFailedResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskFailedScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskFailed(ctx, request, opts...)
}

func (c *metricClient) RespondActivityTaskFailedById(
	ctx context.Context,
	request *workflowservice.RespondActivityTaskFailedByIdRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondActivityTaskFailedByIdResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondActivityTaskFailedByIdScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondActivityTaskFailedById(ctx, request, opts...)
}

func (c *metricClient) RespondQueryTaskCompleted(
	ctx context.Context,
	request *workflowservice.RespondQueryTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondQueryTaskCompletedResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondQueryTaskCompletedScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondQueryTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) RespondWorkflowTaskCompleted(
	ctx context.Context,
	request *workflowservice.RespondWorkflowTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondWorkflowTaskCompletedResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondWorkflowTaskCompletedScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondWorkflowTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) RespondWorkflowTaskFailed(
	ctx context.Context,
	request *workflowservice.RespondWorkflowTaskFailedRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.RespondWorkflowTaskFailedResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientRespondWorkflowTaskFailedScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.RespondWorkflowTaskFailed(ctx, request, opts...)
}

func (c *metricClient) ScanWorkflowExecutions(
	ctx context.Context,
	request *workflowservice.ScanWorkflowExecutionsRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.ScanWorkflowExecutionsResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientScanWorkflowExecutionsScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.ScanWorkflowExecutions(ctx, request, opts...)
}

func (c *metricClient) SignalWithStartWorkflowExecution(
	ctx context.Context,
	request *workflowservice.SignalWithStartWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.SignalWithStartWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientSignalWithStartWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.SignalWithStartWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) SignalWorkflowExecution(
	ctx context.Context,
	request *workflowservice.SignalWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.SignalWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientSignalWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.SignalWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) StartBatchOperation(
	ctx context.Context,
	request *workflowservice.StartBatchOperationRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.StartBatchOperationResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientStartBatchOperationScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.StartBatchOperation(ctx, request, opts...)
}

func (c *metricClient) StartWorkflowExecution(
	ctx context.Context,
	request *workflowservice.StartWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.StartWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientStartWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.StartWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) StopBatchOperation(
	ctx context.Context,
	request *workflowservice.StopBatchOperationRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.StopBatchOperationResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientStopBatchOperationScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.StopBatchOperation(ctx, request, opts...)
}

func (c *metricClient) TerminateWorkflowExecution(
	ctx context.Context,
	request *workflowservice.TerminateWorkflowExecutionRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.TerminateWorkflowExecutionResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientTerminateWorkflowExecutionScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.TerminateWorkflowExecution(ctx, request, opts...)
}

func (c *metricClient) UpdateNamespace(
	ctx context.Context,
	request *workflowservice.UpdateNamespaceRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.UpdateNamespaceResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientUpdateNamespaceScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.UpdateNamespace(ctx, request, opts...)
}

func (c *metricClient) UpdateSchedule(
	ctx context.Context,
	request *workflowservice.UpdateScheduleRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.UpdateScheduleResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientUpdateScheduleScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.UpdateSchedule(ctx, request, opts...)
}

func (c *metricClient) UpdateWorkerBuildIdOrdering(
	ctx context.Context,
	request *workflowservice.UpdateWorkerBuildIdOrderingRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.UpdateWorkerBuildIdOrderingResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientUpdateWorkerBuildIdOrderingScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.UpdateWorkerBuildIdOrdering(ctx, request, opts...)
}

func (c *metricClient) UpdateWorkflow(
	ctx context.Context,
	request *workflowservice.UpdateWorkflowRequest,
	opts ...grpc.CallOption,
) (_ *workflowservice.UpdateWorkflowResponse, retError error) {

	scope, stopwatch := c.startMetricsRecording(metrics.FrontendClientUpdateWorkflowScope)
	defer func() {
		c.finishMetricsRecording(scope, stopwatch, retError)
	}()

	return c.client.UpdateWorkflow(ctx, request, opts...)
}
