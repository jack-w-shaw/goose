// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-goose/goose/v3/identity (interfaces: Authenticator)

// Package mocks is a generated GoMock package.
package mocks

import (
	identity "github.com/go-goose/goose/v3/identity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthenticator is a mock of Authenticator interface
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Auth mocks base method
func (m *MockAuthenticator) Auth(arg0 *identity.Credentials) (*identity.AuthDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", arg0)
	ret0, _ := ret[0].(*identity.AuthDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth
func (mr *MockAuthenticatorMockRecorder) Auth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockAuthenticator)(nil).Auth), arg0)
}
