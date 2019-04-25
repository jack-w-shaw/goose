// Code generated by MockGen. DO NOT EDIT.
// Source: gopkg.in/goose.v2/http (interfaces: HttpClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "gopkg.in/goose.v2/http"
	logging "gopkg.in/goose.v2/logging"
	io "io"
	http0 "net/http"
	url "net/url"
	reflect "reflect"
)

// MockHttpClient is a mock of HttpClient interface
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientMockRecorder
}

// MockHttpClientMockRecorder is the mock recorder for MockHttpClient
type MockHttpClientMockRecorder struct {
	mock *MockHttpClient
}

// NewMockHttpClient creates a new mock instance
func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &MockHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpClient) EXPECT() *MockHttpClientMockRecorder {
	return m.recorder
}

// BinaryRequest mocks base method
func (m *MockHttpClient) BinaryRequest(arg0, arg1, arg2 string, arg3 *http.RequestData, arg4 logging.CompatLogger) error {
	ret := m.ctrl.Call(m, "BinaryRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// BinaryRequest indicates an expected call of BinaryRequest
func (mr *MockHttpClientMockRecorder) BinaryRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BinaryRequest", reflect.TypeOf((*MockHttpClient)(nil).BinaryRequest), arg0, arg1, arg2, arg3, arg4)
}

// Do mocks base method
func (m *MockHttpClient) Do(arg0 *http0.Request) (*http0.Response, error) {
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockHttpClientMockRecorder) Do(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpClient)(nil).Do), arg0)
}

// Get mocks base method
func (m *MockHttpClient) Get(arg0 string) (*http0.Response, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHttpClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClient)(nil).Get), arg0)
}

// Head mocks base method
func (m *MockHttpClient) Head(arg0 string) (*http0.Response, error) {
	ret := m.ctrl.Call(m, "Head", arg0)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head
func (mr *MockHttpClientMockRecorder) Head(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockHttpClient)(nil).Head), arg0)
}

// JsonRequest mocks base method
func (m *MockHttpClient) JsonRequest(arg0, arg1, arg2 string, arg3 *http.RequestData, arg4 logging.CompatLogger) error {
	ret := m.ctrl.Call(m, "JsonRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// JsonRequest indicates an expected call of JsonRequest
func (mr *MockHttpClientMockRecorder) JsonRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JsonRequest", reflect.TypeOf((*MockHttpClient)(nil).JsonRequest), arg0, arg1, arg2, arg3, arg4)
}

// Post mocks base method
func (m *MockHttpClient) Post(arg0, arg1 string, arg2 io.Reader) (*http0.Response, error) {
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post
func (mr *MockHttpClientMockRecorder) Post(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClient)(nil).Post), arg0, arg1, arg2)
}

// PostForm mocks base method
func (m *MockHttpClient) PostForm(arg0 string, arg1 url.Values) (*http0.Response, error) {
	ret := m.ctrl.Call(m, "PostForm", arg0, arg1)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostForm indicates an expected call of PostForm
func (mr *MockHttpClientMockRecorder) PostForm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockHttpClient)(nil).PostForm), arg0, arg1)
}
