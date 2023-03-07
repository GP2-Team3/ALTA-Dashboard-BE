// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mentee "alta-dashboard-be/features/mentee"

	mock "github.com/stretchr/testify/mock"
)

// MenteeData_ is an autogenerated mock type for the MenteeData_ type
type MenteeData_ struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MenteeData_) Create(_a0 mentee.MenteeCore) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentee.MenteeCore) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *MenteeData_) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: page, limit
func (_m *MenteeData_) GetAll(page int, limit int) ([]mentee.MenteeCore, error) {
	ret := _m.Called(page, limit)

	var r0 []mentee.MenteeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]mentee.MenteeCore, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []mentee.MenteeCore); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.MenteeCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: id
func (_m *MenteeData_) GetOne(id int) (mentee.MenteeCore, error) {
	ret := _m.Called(id)

	var r0 mentee.MenteeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (mentee.MenteeCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) mentee.MenteeCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentee.MenteeCore)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, _a1
func (_m *MenteeData_) Update(id int, _a1 mentee.MenteeCore) error {
	ret := _m.Called(id, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, mentee.MenteeCore) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMenteeData_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenteeData_ creates a new instance of MenteeData_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenteeData_(t mockConstructorTestingTNewMenteeData_) *MenteeData_ {
	mock := &MenteeData_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
