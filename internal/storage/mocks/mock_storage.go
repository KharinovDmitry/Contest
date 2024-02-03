// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	domain "contest/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTestRepository is a mock of TestRepository interface.
type MockTestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTestRepositoryMockRecorder
}

// MockTestRepositoryMockRecorder is the mock recorder for MockTestRepository.
type MockTestRepositoryMockRecorder struct {
	mock *MockTestRepository
}

// NewMockTestRepository creates a new mock instance.
func NewMockTestRepository(ctrl *gomock.Controller) *MockTestRepository {
	mock := &MockTestRepository{ctrl: ctrl}
	mock.recorder = &MockTestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestRepository) EXPECT() *MockTestRepositoryMockRecorder {
	return m.recorder
}

// AddTest mocks base method.
func (m *MockTestRepository) AddTest(taskID int, input, expectedResult string, points int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTest", taskID, input, expectedResult, points)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTest indicates an expected call of AddTest.
func (mr *MockTestRepositoryMockRecorder) AddTest(taskID, input, expectedResult, points interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTest", reflect.TypeOf((*MockTestRepository)(nil).AddTest), taskID, input, expectedResult, points)
}

// DeleteTest mocks base method.
func (m *MockTestRepository) DeleteTest(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTest", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTest indicates an expected call of DeleteTest.
func (mr *MockTestRepositoryMockRecorder) DeleteTest(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTest", reflect.TypeOf((*MockTestRepository)(nil).DeleteTest), id)
}

// FindTestByCondition mocks base method.
func (m *MockTestRepository) FindTestByCondition(condition func(domain.Test) bool) (domain.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTestByCondition", condition)
	ret0, _ := ret[0].(domain.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTestByCondition indicates an expected call of FindTestByCondition.
func (mr *MockTestRepositoryMockRecorder) FindTestByCondition(condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTestByCondition", reflect.TypeOf((*MockTestRepository)(nil).FindTestByCondition), condition)
}

// FindTestByID mocks base method.
func (m *MockTestRepository) FindTestByID(id int) (domain.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTestByID", id)
	ret0, _ := ret[0].(domain.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTestByID indicates an expected call of FindTestByID.
func (mr *MockTestRepositoryMockRecorder) FindTestByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTestByID", reflect.TypeOf((*MockTestRepository)(nil).FindTestByID), id)
}

// FindTestsByCondition mocks base method.
func (m *MockTestRepository) FindTestsByCondition(condition func(domain.Test) bool) ([]domain.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTestsByCondition", condition)
	ret0, _ := ret[0].([]domain.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTestsByCondition indicates an expected call of FindTestsByCondition.
func (mr *MockTestRepositoryMockRecorder) FindTestsByCondition(condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTestsByCondition", reflect.TypeOf((*MockTestRepository)(nil).FindTestsByCondition), condition)
}

// FindTestsByTaskID mocks base method.
func (m *MockTestRepository) FindTestsByTaskID(taskID int) ([]domain.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTestsByTaskID", taskID)
	ret0, _ := ret[0].([]domain.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTestsByTaskID indicates an expected call of FindTestsByTaskID.
func (mr *MockTestRepositoryMockRecorder) FindTestsByTaskID(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTestsByTaskID", reflect.TypeOf((*MockTestRepository)(nil).FindTestsByTaskID), taskID)
}

// GetTests mocks base method.
func (m *MockTestRepository) GetTests() ([]domain.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTests")
	ret0, _ := ret[0].([]domain.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTests indicates an expected call of GetTests.
func (mr *MockTestRepositoryMockRecorder) GetTests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTests", reflect.TypeOf((*MockTestRepository)(nil).GetTests))
}

// UpdateTest mocks base method.
func (m *MockTestRepository) UpdateTest(id int, newItem domain.Test) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTest", id, newItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTest indicates an expected call of UpdateTest.
func (mr *MockTestRepositoryMockRecorder) UpdateTest(id, newItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTest", reflect.TypeOf((*MockTestRepository)(nil).UpdateTest), id, newItem)
}