// Code generated by MockGen. DO NOT EDIT.
// Source: state_rebuilder.go
//
// Generated by this command:
//
//	mockgen -package ndc -source state_rebuilder.go -destination state_rebuilder_mock.go
//

// Package ndc is a generated GoMock package.
package ndc

import (
	context "context"
	reflect "reflect"
	time "time"

	persistence "go.temporal.io/server/api/persistence/v1"
	definition "go.temporal.io/server/common/definition"
	interfaces "go.temporal.io/server/service/history/interfaces"
	gomock "go.uber.org/mock/gomock"
)

// MockStateRebuilder is a mock of StateRebuilder interface.
type MockStateRebuilder struct {
	ctrl     *gomock.Controller
	recorder *MockStateRebuilderMockRecorder
	isgomock struct{}
}

// MockStateRebuilderMockRecorder is the mock recorder for MockStateRebuilder.
type MockStateRebuilderMockRecorder struct {
	mock *MockStateRebuilder
}

// NewMockStateRebuilder creates a new mock instance.
func NewMockStateRebuilder(ctrl *gomock.Controller) *MockStateRebuilder {
	mock := &MockStateRebuilder{ctrl: ctrl}
	mock.recorder = &MockStateRebuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateRebuilder) EXPECT() *MockStateRebuilderMockRecorder {
	return m.recorder
}

// Rebuild mocks base method.
func (m *MockStateRebuilder) Rebuild(ctx context.Context, now time.Time, baseWorkflowIdentifier definition.WorkflowKey, baseBranchToken []byte, baseLastEventID int64, baseLastEventVersion *int64, targetWorkflowIdentifier definition.WorkflowKey, targetBranchToken []byte, requestID string) (interfaces.MutableState, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebuild", ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID)
	ret0, _ := ret[0].(interfaces.MutableState)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Rebuild indicates an expected call of Rebuild.
func (mr *MockStateRebuilderMockRecorder) Rebuild(ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebuild", reflect.TypeOf((*MockStateRebuilder)(nil).Rebuild), ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID)
}

// RebuildWithCurrentMutableState mocks base method.
func (m *MockStateRebuilder) RebuildWithCurrentMutableState(ctx context.Context, now time.Time, baseWorkflowIdentifier definition.WorkflowKey, baseBranchToken []byte, baseLastEventID int64, baseLastEventVersion *int64, targetWorkflowIdentifier definition.WorkflowKey, targetBranchToken []byte, requestID string, currentMutableState *persistence.WorkflowMutableState) (interfaces.MutableState, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebuildWithCurrentMutableState", ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID, currentMutableState)
	ret0, _ := ret[0].(interfaces.MutableState)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RebuildWithCurrentMutableState indicates an expected call of RebuildWithCurrentMutableState.
func (mr *MockStateRebuilderMockRecorder) RebuildWithCurrentMutableState(ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID, currentMutableState any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebuildWithCurrentMutableState", reflect.TypeOf((*MockStateRebuilder)(nil).RebuildWithCurrentMutableState), ctx, now, baseWorkflowIdentifier, baseBranchToken, baseLastEventID, baseLastEventVersion, targetWorkflowIdentifier, targetBranchToken, requestID, currentMutableState)
}
