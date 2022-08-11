// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/cli/pkg/plugin/registry (interfaces: Registry)

// Package registry is a generated GoMock package.
package registry

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// CheckUpdate mocks base method.
func (m *MockRegistry) CheckUpdate(arg0 context.Context, arg1 Provider) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUpdate", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUpdate indicates an expected call of CheckUpdate.
func (mr *MockRegistryMockRecorder) CheckUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpdate", reflect.TypeOf((*MockRegistry)(nil).CheckUpdate), arg0, arg1)
}

// Download mocks base method.
func (m *MockRegistry) Download(arg0 context.Context, arg1 Provider, arg2 bool) (ProviderBinary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2)
	ret0, _ := ret[0].(ProviderBinary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockRegistryMockRecorder) Download(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockRegistry)(nil).Download), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockRegistry) Get(arg0, arg1 string) (ProviderBinary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(ProviderBinary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRegistryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRegistry)(nil).Get), arg0, arg1)
}
