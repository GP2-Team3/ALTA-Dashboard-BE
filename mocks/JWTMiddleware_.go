// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// JWTMiddleware_ is an autogenerated mock type for the JWTMiddleware_ type
type JWTMiddleware_ struct {
	mock.Mock
}

// ExtractToken provides a mock function with given fields: e
func (_m *JWTMiddleware_) ExtractToken(e echo.Context) (uint, string, error) {
	ret := _m.Called(e)

	var r0 uint
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context) (uint, string, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(echo.Context) string); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(echo.Context) error); ok {
		r2 = rf(e)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewJWTMiddleware_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewJWTMiddleware_ creates a new instance of JWTMiddleware_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJWTMiddleware_(t mockConstructorTestingTNewJWTMiddleware_) *JWTMiddleware_ {
	mock := &JWTMiddleware_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
