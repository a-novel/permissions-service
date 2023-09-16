// Code generated by mockery v2.33.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/a-novel/authorizations-service/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GetUserScopesService is an autogenerated mock type for the GetUserScopesService type
type GetUserScopesService struct {
	mock.Mock
}

type GetUserScopesService_Expecter struct {
	mock *mock.Mock
}

func (_m *GetUserScopesService) EXPECT() *GetUserScopesService_Expecter {
	return &GetUserScopesService_Expecter{mock: &_m.Mock}
}

// GetUserScopes provides a mock function with given fields: ctx, tokenRaw
func (_m *GetUserScopesService) GetUserScopes(ctx context.Context, tokenRaw string) (models.Scopes, error) {
	ret := _m.Called(ctx, tokenRaw)

	var r0 models.Scopes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Scopes, error)); ok {
		return rf(ctx, tokenRaw)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Scopes); ok {
		r0 = rf(ctx, tokenRaw)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Scopes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenRaw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserScopesService_GetUserScopes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserScopes'
type GetUserScopesService_GetUserScopes_Call struct {
	*mock.Call
}

// GetUserScopes is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenRaw string
func (_e *GetUserScopesService_Expecter) GetUserScopes(ctx interface{}, tokenRaw interface{}) *GetUserScopesService_GetUserScopes_Call {
	return &GetUserScopesService_GetUserScopes_Call{Call: _e.mock.On("GetUserScopes", ctx, tokenRaw)}
}

func (_c *GetUserScopesService_GetUserScopes_Call) Run(run func(ctx context.Context, tokenRaw string)) *GetUserScopesService_GetUserScopes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *GetUserScopesService_GetUserScopes_Call) Return(_a0 models.Scopes, _a1 error) *GetUserScopesService_GetUserScopes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetUserScopesService_GetUserScopes_Call) RunAndReturn(run func(context.Context, string) (models.Scopes, error)) *GetUserScopesService_GetUserScopes_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetUserScopesService creates a new instance of GetUserScopesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetUserScopesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetUserScopesService {
	mock := &GetUserScopesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
