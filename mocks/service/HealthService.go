// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HealthService is an autogenerated mock type for the HealthService type
type HealthService struct {
	mock.Mock
}

// Check provides a mock function with given fields:
func (_m *HealthService) Check() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
