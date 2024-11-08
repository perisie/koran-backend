// Code generated by MockGen. DO NOT EDIT.
// Source: ./dao.go

// Package favorite is a generated GoMock package.
package favorite

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "perisie.com/koran/models"
)

// FavDaoMock is a mock of FavDao interface.
type FavDaoMock struct {
	ctrl     *gomock.Controller
	recorder *FavDaoMockMockRecorder
}

// FavDaoMockMockRecorder is the mock recorder for FavDaoMock.
type FavDaoMockMockRecorder struct {
	mock *FavDaoMock
}

// NewFavDaoMock creates a new mock instance.
func NewFavDaoMock(ctrl *gomock.Controller) *FavDaoMock {
	mock := &FavDaoMock{ctrl: ctrl}
	mock.recorder = &FavDaoMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FavDaoMock) EXPECT() *FavDaoMockMockRecorder {
	return m.recorder
}

// AddFavVerse mocks base method.
func (m *FavDaoMock) AddFavVerse(email string, surah, verse int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavVerse", email, surah, verse)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavVerse indicates an expected call of AddFavVerse.
func (mr *FavDaoMockMockRecorder) AddFavVerse(email, surah, verse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavVerse", reflect.TypeOf((*FavDaoMock)(nil).AddFavVerse), email, surah, verse)
}

// DeleteFav mocks base method.
func (m *FavDaoMock) DeleteFav(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFav", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFav indicates an expected call of DeleteFav.
func (mr *FavDaoMockMockRecorder) DeleteFav(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFav", reflect.TypeOf((*FavDaoMock)(nil).DeleteFav), id)
}

// QueryUserFavsByEmail mocks base method.
func (m *FavDaoMock) QueryUserFavsByEmail(email string) ([]*models.Fav, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserFavsByEmail", email)
	ret0, _ := ret[0].([]*models.Fav)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserFavsByEmail indicates an expected call of QueryUserFavsByEmail.
func (mr *FavDaoMockMockRecorder) QueryUserFavsByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserFavsByEmail", reflect.TypeOf((*FavDaoMock)(nil).QueryUserFavsByEmail), email)
}
