// Code generated by mockery v2.33.2. DO NOT EDIT.

package daomocks

import (
	context "context"

	dao "github.com/a-novel/permissions-service/pkg/dao"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// UserPermissionsRepository is an autogenerated mock type for the UserPermissionsRepository type
type UserPermissionsRepository struct {
	mock.Mock
}

type UserPermissionsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserPermissionsRepository) EXPECT() *UserPermissionsRepository_Expecter {
	return &UserPermissionsRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, userID
func (_m *UserPermissionsRepository) Get(ctx context.Context, userID uuid.UUID) (*dao.UserPermissions, error) {
	ret := _m.Called(ctx, userID)

	var r0 *dao.UserPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*dao.UserPermissions, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *dao.UserPermissions); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.UserPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserPermissionsRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type UserPermissionsRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *UserPermissionsRepository_Expecter) Get(ctx interface{}, userID interface{}) *UserPermissionsRepository_Get_Call {
	return &UserPermissionsRepository_Get_Call{Call: _e.mock.On("Get", ctx, userID)}
}

func (_c *UserPermissionsRepository_Get_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *UserPermissionsRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *UserPermissionsRepository_Get_Call) Return(_a0 *dao.UserPermissions, _a1 error) *UserPermissionsRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserPermissionsRepository_Get_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*dao.UserPermissions, error)) *UserPermissionsRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, userID, core, fields, now
func (_m *UserPermissionsRepository) Set(ctx context.Context, userID uuid.UUID, core dao.UserPermissionsCore, fields dao.PermissionsFields, now time.Time) (*dao.UserPermissions, error) {
	ret := _m.Called(ctx, userID, core, fields, now)

	var r0 *dao.UserPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dao.UserPermissionsCore, dao.PermissionsFields, time.Time) (*dao.UserPermissions, error)); ok {
		return rf(ctx, userID, core, fields, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dao.UserPermissionsCore, dao.PermissionsFields, time.Time) *dao.UserPermissions); ok {
		r0 = rf(ctx, userID, core, fields, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.UserPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, dao.UserPermissionsCore, dao.PermissionsFields, time.Time) error); ok {
		r1 = rf(ctx, userID, core, fields, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserPermissionsRepository_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type UserPermissionsRepository_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - core dao.UserPermissionsCore
//   - fields dao.PermissionsFields
//   - now time.Time
func (_e *UserPermissionsRepository_Expecter) Set(ctx interface{}, userID interface{}, core interface{}, fields interface{}, now interface{}) *UserPermissionsRepository_Set_Call {
	return &UserPermissionsRepository_Set_Call{Call: _e.mock.On("Set", ctx, userID, core, fields, now)}
}

func (_c *UserPermissionsRepository_Set_Call) Run(run func(ctx context.Context, userID uuid.UUID, core dao.UserPermissionsCore, fields dao.PermissionsFields, now time.Time)) *UserPermissionsRepository_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(dao.UserPermissionsCore), args[3].(dao.PermissionsFields), args[4].(time.Time))
	})
	return _c
}

func (_c *UserPermissionsRepository_Set_Call) Return(_a0 *dao.UserPermissions, _a1 error) *UserPermissionsRepository_Set_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserPermissionsRepository_Set_Call) RunAndReturn(run func(context.Context, uuid.UUID, dao.UserPermissionsCore, dao.PermissionsFields, time.Time) (*dao.UserPermissions, error)) *UserPermissionsRepository_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserPermissionsRepository creates a new instance of UserPermissionsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserPermissionsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserPermissionsRepository {
	mock := &UserPermissionsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
