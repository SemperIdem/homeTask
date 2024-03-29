// Code generated by MockGen. DO NOT EDIT.
// Source: homeTask/server (interfaces: NumbersFetcher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNumbersFetcher is a mock of NumbersFetcher interface.
type MockNumbersFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockNumbersFetcherMockRecorder
}

// MockNumbersFetcherMockRecorder is the mock recorder for MockNumbersFetcher.
type MockNumbersFetcherMockRecorder struct {
	mock *MockNumbersFetcher
}

// NewMockNumbersFetcher creates a new mock instance.
func NewMockNumbersFetcher(ctrl *gomock.Controller) *MockNumbersFetcher {
	mock := &MockNumbersFetcher{ctrl: ctrl}
	mock.recorder = &MockNumbersFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNumbersFetcher) EXPECT() *MockNumbersFetcherMockRecorder {
	return m.recorder
}

// ProcessUrls mocks base method.
func (m *MockNumbersFetcher) ProcessUrls(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessUrls", arg0)
}

// ProcessUrls indicates an expected call of ProcessUrls.
func (mr *MockNumbersFetcherMockRecorder) ProcessUrls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessUrls", reflect.TypeOf((*MockNumbersFetcher)(nil).ProcessUrls), arg0)
}

// Receive mocks base method.
func (m *MockNumbersFetcher) Receive() chan []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(chan []int)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *MockNumbersFetcherMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockNumbersFetcher)(nil).Receive))
}
