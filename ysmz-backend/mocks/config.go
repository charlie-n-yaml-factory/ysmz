// Code generated by MockGen. DO NOT EDIT.
// Source: ./config/default.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	config "github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/config"
	gomock "github.com/golang/mock/gomock"
)

// MockConfigInterface is a mock of ConfigInterface interface.
type MockConfigInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigInterfaceMockRecorder
}

// MockConfigInterfaceMockRecorder is the mock recorder for MockConfigInterface.
type MockConfigInterfaceMockRecorder struct {
	mock *MockConfigInterface
}

// NewMockConfigInterface creates a new mock instance.
func NewMockConfigInterface(ctrl *gomock.Controller) *MockConfigInterface {
	mock := &MockConfigInterface{ctrl: ctrl}
	mock.recorder = &MockConfigInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigInterface) EXPECT() *MockConfigInterfaceMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockConfigInterface) Config() *config.ConfigStruct {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.ConfigStruct)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockConfigInterfaceMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfigInterface)(nil).Config))
}
