// Code generated by mockery v2.33.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// SetUserAuthorizationService is an autogenerated mock type for the SetUserAuthorizationService type
type SetUserAuthorizationService struct {
	mock.Mock
}

type SetUserAuthorizationService_Expecter struct {
	mock *mock.Mock
}

func (_m *SetUserAuthorizationService) EXPECT() *SetUserAuthorizationService_Expecter {
	return &SetUserAuthorizationService_Expecter{mock: &_m.Mock}
}

// Set provides a mock function with given fields: ctx, userID, setFields, unsetFields, now
func (_m *SetUserAuthorizationService) Set(ctx context.Context, userID uuid.UUID, setFields []string, unsetFields []string, now time.Time) error {
	ret := _m.Called(ctx, userID, setFields, unsetFields, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []string, []string, time.Time) error); ok {
		r0 = rf(ctx, userID, setFields, unsetFields, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUserAuthorizationService_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type SetUserAuthorizationService_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - setFields []string
//   - unsetFields []string
//   - now time.Time
func (_e *SetUserAuthorizationService_Expecter) Set(ctx interface{}, userID interface{}, setFields interface{}, unsetFields interface{}, now interface{}) *SetUserAuthorizationService_Set_Call {
	return &SetUserAuthorizationService_Set_Call{Call: _e.mock.On("Set", ctx, userID, setFields, unsetFields, now)}
}

func (_c *SetUserAuthorizationService_Set_Call) Run(run func(ctx context.Context, userID uuid.UUID, setFields []string, unsetFields []string, now time.Time)) *SetUserAuthorizationService_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]string), args[3].([]string), args[4].(time.Time))
	})
	return _c
}

func (_c *SetUserAuthorizationService_Set_Call) Return(_a0 error) *SetUserAuthorizationService_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SetUserAuthorizationService_Set_Call) RunAndReturn(run func(context.Context, uuid.UUID, []string, []string, time.Time) error) *SetUserAuthorizationService_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewSetUserAuthorizationService creates a new instance of SetUserAuthorizationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSetUserAuthorizationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SetUserAuthorizationService {
	mock := &SetUserAuthorizationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
