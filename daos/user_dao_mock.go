// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_dao.go

// Package daos is a generated GoMock package.
package daos

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	beans "perisie.com/koran/beans"
)

// UserDaoMock is a mock of UserDao interface.
type UserDaoMock struct {
	ctrl     *gomock.Controller
	recorder *UserDaoMockMockRecorder
}

// UserDaoMockMockRecorder is the mock recorder for UserDaoMock.
type UserDaoMockMockRecorder struct {
	mock *UserDaoMock
}

// NewUserDaoMock creates a new mock instance.
func NewUserDaoMock(ctrl *gomock.Controller) *UserDaoMock {
	mock := &UserDaoMock{ctrl: ctrl}
	mock.recorder = &UserDaoMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserDaoMock) EXPECT() *UserDaoMockMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *UserDaoMock) CreateUser(email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *UserDaoMockMockRecorder) CreateUser(email, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*UserDaoMock)(nil).CreateUser), email, token)
}

// QueryUserByEmail mocks base method.
func (m *UserDaoMock) QueryUserByEmail(email string) (*beans.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserByEmail", email)
	ret0, _ := ret[0].(*beans.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserByEmail indicates an expected call of QueryUserByEmail.
func (mr *UserDaoMockMockRecorder) QueryUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserByEmail", reflect.TypeOf((*UserDaoMock)(nil).QueryUserByEmail), email)
}

// QueryUserByToken mocks base method.
func (m *UserDaoMock) QueryUserByToken(token string) (*beans.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserByToken", token)
	ret0, _ := ret[0].(*beans.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserByToken indicates an expected call of QueryUserByToken.
func (mr *UserDaoMockMockRecorder) QueryUserByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserByToken", reflect.TypeOf((*UserDaoMock)(nil).QueryUserByToken), token)
}

// UpdateUserCurrentPointer mocks base method.
func (m *UserDaoMock) UpdateUserCurrentPointer(email, currentPointer string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserCurrentPointer", email, currentPointer)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserCurrentPointer indicates an expected call of UpdateUserCurrentPointer.
func (mr *UserDaoMockMockRecorder) UpdateUserCurrentPointer(email, currentPointer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserCurrentPointer", reflect.TypeOf((*UserDaoMock)(nil).UpdateUserCurrentPointer), email, currentPointer)
}

// UpdateUserToken mocks base method.
func (m *UserDaoMock) UpdateUserToken(email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserToken", email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserToken indicates an expected call of UpdateUserToken.
func (mr *UserDaoMockMockRecorder) UpdateUserToken(email, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserToken", reflect.TypeOf((*UserDaoMock)(nil).UpdateUserToken), email, token)
}
