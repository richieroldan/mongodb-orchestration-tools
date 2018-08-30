// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import api "github.com/percona/dcos-mongo-tools/internal/api"
import mock "github.com/stretchr/testify/mock"

// PodTask is an autogenerated mock type for the PodTask type
type PodTask struct {
	mock.Mock
}

// GetEnvVar provides a mock function with given fields: variableName
func (_m *PodTask) GetEnvVar(variableName string) (string, error) {
	ret := _m.Called(variableName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(variableName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(variableName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMongoHostname provides a mock function with given fields: frameworkName
func (_m *PodTask) GetMongoHostname(frameworkName string) string {
	ret := _m.Called(frameworkName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(frameworkName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMongoPort provides a mock function with given fields:
func (_m *PodTask) GetMongoPort() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMongoReplsetName provides a mock function with given fields:
func (_m *PodTask) GetMongoReplsetName() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasState provides a mock function with given fields:
func (_m *PodTask) HasState() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsMongodTask provides a mock function with given fields:
func (_m *PodTask) IsMongodTask() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsMongosTask provides a mock function with given fields:
func (_m *PodTask) IsMongosTask() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRemovedMongod provides a mock function with given fields:
func (_m *PodTask) IsRemovedMongod() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRunning provides a mock function with given fields:
func (_m *PodTask) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *PodTask) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// State provides a mock function with given fields:
func (_m *PodTask) State() api.PodTaskState {
	ret := _m.Called()

	var r0 api.PodTaskState
	if rf, ok := ret.Get(0).(func() api.PodTaskState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.PodTaskState)
	}

	return r0
}