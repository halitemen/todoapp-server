// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\P1561\Desktop\gitlab\todoapp-server\service\todo_service.go

// Package mock_service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockService) CreateTodo(i interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", i)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockServiceMockRecorder) CreateTodo(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockService)(nil).CreateTodo), i)
}

// GetTodos mocks base method.
func (m *MockService) GetTodos() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodos")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockServiceMockRecorder) GetTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockService)(nil).GetTodos))
}