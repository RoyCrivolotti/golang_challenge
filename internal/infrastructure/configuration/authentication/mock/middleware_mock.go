// Code generated by MockGen. DO NOT EDIT.
// Source: ./middleware.go

// Package mock_authentication is a generated GoMock package.
package mock_authentication

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAuthenticationMiddleware is a mock of IAuthenticationMiddleware interface.
type MockIAuthenticationMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthenticationMiddlewareMockRecorder
}

// MockIAuthenticationMiddlewareMockRecorder is the mock recorder for MockIAuthenticationMiddleware.
type MockIAuthenticationMiddlewareMockRecorder struct {
	mock *MockIAuthenticationMiddleware
}

// NewMockIAuthenticationMiddleware creates a new mock instance.
func NewMockIAuthenticationMiddleware(ctrl *gomock.Controller) *MockIAuthenticationMiddleware {
	mock := &MockIAuthenticationMiddleware{ctrl: ctrl}
	mock.recorder = &MockIAuthenticationMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthenticationMiddleware) EXPECT() *MockIAuthenticationMiddlewareMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockIAuthenticationMiddleware) Authenticate(next http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", next)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockIAuthenticationMiddlewareMockRecorder) Authenticate(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockIAuthenticationMiddleware)(nil).Authenticate), next)
}
