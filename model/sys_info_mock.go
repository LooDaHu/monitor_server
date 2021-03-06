// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/sys_info.go

// Package model is a generated GoMock package.
package model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
)

// MockSysInfoer is a mock of SysInfoer interface.
type MockSysInfoer struct {
	ctrl     *gomock.Controller
	recorder *MockSysInfoerMockRecorder
}

// MockSysInfoerMockRecorder is the mock recorder for MockSysInfoer.
type MockSysInfoerMockRecorder struct {
	mock *MockSysInfoer
}

// NewMockSysInfoer creates a new mock instance.
func NewMockSysInfoer(ctrl *gomock.Controller) *MockSysInfoer {
	mock := &MockSysInfoer{ctrl: ctrl}
	mock.recorder = &MockSysInfoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSysInfoer) EXPECT() *MockSysInfoerMockRecorder {
	return m.recorder
}

// CreateSystemInfo mocks base method.
func (m *MockSysInfoer) CreateSystemInfo(sysInfo *SysInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSystemInfo", sysInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSystemInfo indicates an expected call of CreateSystemInfo.
func (mr *MockSysInfoerMockRecorder) CreateSystemInfo(sysInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSystemInfo", reflect.TypeOf((*MockSysInfoer)(nil).CreateSystemInfo), sysInfo)
}

// DeleteSystemInfo mocks base method.
func (m *MockSysInfoer) DeleteSystemInfo(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSystemInfo", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSystemInfo indicates an expected call of DeleteSystemInfo.
func (mr *MockSysInfoerMockRecorder) DeleteSystemInfo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSystemInfo", reflect.TypeOf((*MockSysInfoer)(nil).DeleteSystemInfo), id)
}

// RetrieveSystemInfo mocks base method.
func (m *MockSysInfoer) RetrieveSystemInfo(filter bson.M) (*SysInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveSystemInfo", filter)
	ret0, _ := ret[0].(*SysInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveSystemInfo indicates an expected call of RetrieveSystemInfo.
func (mr *MockSysInfoerMockRecorder) RetrieveSystemInfo(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveSystemInfo", reflect.TypeOf((*MockSysInfoer)(nil).RetrieveSystemInfo), filter)
}

// UpdateSystemInfo mocks base method.
func (m *MockSysInfoer) UpdateSystemInfo(sysInfo *SysInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSystemInfo", sysInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSystemInfo indicates an expected call of UpdateSystemInfo.
func (mr *MockSysInfoerMockRecorder) UpdateSystemInfo(sysInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSystemInfo", reflect.TypeOf((*MockSysInfoer)(nil).UpdateSystemInfo), sysInfo)
}
