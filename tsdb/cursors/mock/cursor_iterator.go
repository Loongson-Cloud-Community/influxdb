// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/tsdb/cursors (interfaces: CursorIterator)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cursors "github.com/influxdata/influxdb/tsdb/cursors"
)

// MockCursorIterator is a mock of CursorIterator interface
type MockCursorIterator struct {
	ctrl     *gomock.Controller
	recorder *MockCursorIteratorMockRecorder
}

// MockCursorIteratorMockRecorder is the mock recorder for MockCursorIterator
type MockCursorIteratorMockRecorder struct {
	mock *MockCursorIterator
}

// NewMockCursorIterator creates a new mock instance
func NewMockCursorIterator(ctrl *gomock.Controller) *MockCursorIterator {
	mock := &MockCursorIterator{ctrl: ctrl}
	mock.recorder = &MockCursorIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCursorIterator) EXPECT() *MockCursorIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockCursorIterator) Next(arg0 context.Context, arg1 *cursors.CursorRequest) (cursors.Cursor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0, arg1)
	ret0, _ := ret[0].(cursors.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockCursorIteratorMockRecorder) Next(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCursorIterator)(nil).Next), arg0, arg1)
}

// Stats mocks base method
func (m *MockCursorIterator) Stats() cursors.CursorStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(cursors.CursorStats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockCursorIteratorMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockCursorIterator)(nil).Stats))
}
