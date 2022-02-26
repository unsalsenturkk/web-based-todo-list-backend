// Code generated by MockGen. DO NOT EDIT.
// Source: .\service\service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	reflect "reflect"
	models "web-based-todo-list-backend/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// AddTodoList mocks base method.
func (m *MockIService) AddTodoList(todo string) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTodoList", todo)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTodoList indicates an expected call of AddTodoList.
func (mr *MockIServiceMockRecorder) AddTodoList(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTodoList", reflect.TypeOf((*MockIService)(nil).AddTodoList), todo)
}

// GetTodoList mocks base method.
func (m *MockIService) GetTodoList() (*models.DataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoList")
	ret0, _ := ret[0].(*models.DataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoList indicates an expected call of GetTodoList.
func (mr *MockIServiceMockRecorder) GetTodoList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoList", reflect.TypeOf((*MockIService)(nil).GetTodoList))
}