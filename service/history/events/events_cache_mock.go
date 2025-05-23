// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -package events -source cache.go -destination events_cache_mock.go
//

// Package events is a generated GoMock package.
package events

import (
	context "context"
	reflect "reflect"

	history "go.temporal.io/api/history/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
	isgomock struct{}
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// DeleteEvent mocks base method.
func (m *MockCache) DeleteEvent(key EventKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteEvent", key)
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockCacheMockRecorder) DeleteEvent(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockCache)(nil).DeleteEvent), key)
}

// GetEvent mocks base method.
func (m *MockCache) GetEvent(ctx context.Context, shardID int32, key EventKey, firstEventID int64, branchToken []byte) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, shardID, key, firstEventID, branchToken)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockCacheMockRecorder) GetEvent(ctx, shardID, key, firstEventID, branchToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockCache)(nil).GetEvent), ctx, shardID, key, firstEventID, branchToken)
}

// PutEvent mocks base method.
func (m *MockCache) PutEvent(key EventKey, event *history.HistoryEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutEvent", key, event)
}

// PutEvent indicates an expected call of PutEvent.
func (mr *MockCacheMockRecorder) PutEvent(key, event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutEvent", reflect.TypeOf((*MockCache)(nil).PutEvent), key, event)
}
