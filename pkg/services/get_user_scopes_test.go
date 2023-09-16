package services_test

import (
	"context"
	"github.com/a-novel/bunovel"
	apiclients "github.com/a-novel/go-api-clients"
	apiclientsmocks "github.com/a-novel/go-api-clients/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/adapters"
	"github.com/a-novel/permissions-service/pkg/dao"
	daomocks "github.com/a-novel/permissions-service/pkg/dao/mocks"
	"github.com/a-novel/permissions-service/pkg/models"
	"github.com/a-novel/permissions-service/pkg/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUserScopesService(t *testing.T) {
	data := []struct {
		name string

		tokenRaw string

		authClientResp *apiclients.UserTokenStatus
		authClientErr  error

		shouldCallDAO bool
		daoResp       *dao.UserPermissions
		daoErr        error

		expect    models.Scopes
		expectErr error
	}{
		{
			name:     "Success",
			tokenRaw: "token",
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallDAO: true,
			daoResp: &dao.UserPermissions{
				UserPermissionsCore: dao.UserPermissionsCore{
					ValidatedAccount: true,
				},
			},
			expect: append(adapters.DefaultScopes, adapters.ValidatedAccountScopes...),
		},
		{
			name:     "Success/NotFound",
			tokenRaw: "token",
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallDAO: true,
			daoErr:        bunovel.ErrNotFound,
		},
		{
			name:     "Error/DAOFailure",
			tokenRaw: "token",
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallDAO: true,
			daoErr:        fooErr,
			expectErr:     fooErr,
		},
		{
			name:           "Error/NotAuthenticated",
			tokenRaw:       "token",
			authClientResp: &apiclients.UserTokenStatus{},
			expectErr:      goframework.ErrInvalidCredentials,
		},
		{
			name:          "Error/AuthClientFailure",
			tokenRaw:      "token",
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewUserPermissionsRepository(t)
			authClient := apiclientsmocks.NewAuthClient(t)

			authClient.On("IntrospectToken", context.Background(), d.tokenRaw).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallDAO {
				repository.On("Get", context.Background(), d.authClientResp.Token.Payload.ID).Return(d.daoResp, d.daoErr)
			}

			service := services.NewGetUserScopesService(repository, authClient)
			res, err := service.GetUserScopes(context.Background(), d.tokenRaw)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, res)

			repository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
