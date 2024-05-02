// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	environment "github.com/jfelipearaujo-org/ms-payment-management/internal/environment"
	mock "github.com/stretchr/testify/mock"
)

// MockEnvironment is an autogenerated mock type for the Environment type
type MockEnvironment struct {
	mock.Mock
}

// GetEnvironment provides a mock function with given fields: ctx
func (_m *MockEnvironment) GetEnvironment(ctx context.Context) (*environment.Config, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEnvironment")
	}

	var r0 *environment.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*environment.Config, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *environment.Config); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*environment.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnvironmentFromFile provides a mock function with given fields: ctx, fileName
func (_m *MockEnvironment) GetEnvironmentFromFile(ctx context.Context, fileName string) (*environment.Config, error) {
	ret := _m.Called(ctx, fileName)

	if len(ret) == 0 {
		panic("no return value specified for GetEnvironmentFromFile")
	}

	var r0 *environment.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*environment.Config, error)); ok {
		return rf(ctx, fileName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *environment.Config); ok {
		r0 = rf(ctx, fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*environment.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEnvironment creates a new instance of MockEnvironment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnvironment(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnvironment {
	mock := &MockEnvironment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
