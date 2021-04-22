// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\P1561\Desktop\gitlab\todoapp-server\repository\repository.go

// Package mock_repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockRepository) CreateTodo(i interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", i)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockRepositoryMockRecorder) CreateTodo(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockRepository)(nil).CreateTodo), i)
}

// GetTodos mocks base method.
func (m *MockRepository) GetTodos() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodos")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockRepositoryMockRecorder) GetTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockRepository)(nil).GetTodos))
}
