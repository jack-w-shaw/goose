// Code generated by MockGen. DO NOT EDIT.
// Source: gopkg.in/goose.v3/logging (interfaces: CompatLogger)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCompatLogger is a mock of CompatLogger interface
type MockCompatLogger struct {
	ctrl     *gomock.Controller
	recorder *MockCompatLoggerMockRecorder
}

// MockCompatLoggerMockRecorder is the mock recorder for MockCompatLogger
type MockCompatLoggerMockRecorder struct {
	mock *MockCompatLogger
}

// NewMockCompatLogger creates a new mock instance
func NewMockCompatLogger(ctrl *gomock.Controller) *MockCompatLogger {
	mock := &MockCompatLogger{ctrl: ctrl}
	mock.recorder = &MockCompatLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompatLogger) EXPECT() *MockCompatLoggerMockRecorder {
	return m.recorder
}

// Printf mocks base method
func (m *MockCompatLogger) Printf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf
func (mr *MockCompatLoggerMockRecorder) Printf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockCompatLogger)(nil).Printf), varargs...)
}
