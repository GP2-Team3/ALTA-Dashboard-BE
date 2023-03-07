// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	logs "alta-dashboard-be/features/logs"

	mock "github.com/stretchr/testify/mock"
)

// LogService_ is an autogenerated mock type for the LogServiceInterface_ type
type LogService_ struct {
	mock.Mock
}

// Create provides a mock function with given fields: logInput, loggedInUserId
func (_m *LogService_) Create(logInput logs.LogEntity, loggedInUserId uint) (logs.LogEntity, error) {
	ret := _m.Called(logInput, loggedInUserId)

	var r0 logs.LogEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(logs.LogEntity, uint) (logs.LogEntity, error)); ok {
		return rf(logInput, loggedInUserId)
	}
	if rf, ok := ret.Get(0).(func(logs.LogEntity, uint) logs.LogEntity); ok {
		r0 = rf(logInput, loggedInUserId)
	} else {
		r0 = ret.Get(0).(logs.LogEntity)
	}

	if rf, ok := ret.Get(1).(func(logs.LogEntity, uint) error); ok {
		r1 = rf(logInput, loggedInUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetData provides a mock function with given fields: searchedMenteeId, limit, offset
func (_m *LogService_) GetData(searchedMenteeId uint, limit int, offset int) (map[string]interface{}, error) {
	ret := _m.Called(searchedMenteeId, limit, offset)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int) (map[string]interface{}, error)); ok {
		return rf(searchedMenteeId, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int) map[string]interface{}); ok {
		r0 = rf(searchedMenteeId, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(searchedMenteeId, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLogService_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogService_ creates a new instance of LogService_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogService_(t mockConstructorTestingTNewLogService_) *LogService_ {
	mock := &LogService_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
