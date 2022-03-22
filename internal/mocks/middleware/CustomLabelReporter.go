// Code generated by mockery v2.8.0. DO NOT EDIT.

package middleware

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CustomLabelReporter is an autogenerated mock type for the CustomLabelReporter type
type CustomLabelReporter struct {
	mock.Mock
}

// BytesWritten provides a mock function with given fields:
func (_m *CustomLabelReporter) BytesWritten() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *CustomLabelReporter) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// CustomLabels provides a mock function with given fields:
func (_m *CustomLabelReporter) CustomLabels() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *CustomLabelReporter) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StatusCode provides a mock function with given fields:
func (_m *CustomLabelReporter) StatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// URLPath provides a mock function with given fields:
func (_m *CustomLabelReporter) URLPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}