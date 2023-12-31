// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/user/repository/user_repository.go
//
// Generated by this command:
//
//	mockgen -source=src/model/user/repository/user_repository.go -destination=src/test/mocks/user_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, *rest_errors.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userDomain)
	ret0, _ := ret[0].(user_model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_errors.RestError)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(userDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), userDomain)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(arg0 string) *rest_errors.RestError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*rest_errors.RestError)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), arg0)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(arg0 string) (user_model.UserDomainInterface, *rest_errors.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(user_model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_errors.RestError)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), arg0)
}

// FindUserByID mocks base method.
func (m *MockUserRepository) FindUserByID(arg0 string) (user_model.UserDomainInterface, *rest_errors.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0)
	ret0, _ := ret[0].(user_model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_errors.RestError)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserRepositoryMockRecorder) FindUserByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByID), arg0)
}

// UpdateUserById mocks base method.
func (m *MockUserRepository) UpdateUserById(arg0 string, arg1 user_model.UserDomainInterface) *rest_errors.RestError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", arg0, arg1)
	ret0, _ := ret[0].(*rest_errors.RestError)
	return ret0
}

// UpdateUserById indicates an expected call of UpdateUserById.
func (mr *MockUserRepositoryMockRecorder) UpdateUserById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserById), arg0, arg1)
}
