package services_test

import (
	"context"
	"github.com/a-novel/authorizations-service/pkg/dao"
	daomocks "github.com/a-novel/authorizations-service/pkg/dao/mocks"
	"github.com/a-novel/authorizations-service/pkg/models"
	"github.com/a-novel/authorizations-service/pkg/services"
	"github.com/a-novel/bunovel"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHasUserScopeService(t *testing.T) {
	data := []struct {
		name string

		userID uuid.UUID
		scope  string

		daoResp *dao.UserAuthorizations
		daoErr  error

		expect    bool
		expectErr error
	}{
		{
			name:   "Success/HasScope",
			userID: goframework.NumberUUID(1),
			scope:  string(models.CanPostImproveSuggestion),
			daoResp: &dao.UserAuthorizations{
				UserAuthorizationsCore: dao.UserAuthorizationsCore{
					ValidatedAccount: true,
				},
			},
			expect: true,
		},
		{
			name:   "Success/HasNotScope",
			userID: goframework.NumberUUID(1),
			scope:  string(models.CanUseOpenAIPlayground),
			daoResp: &dao.UserAuthorizations{
				UserAuthorizationsCore: dao.UserAuthorizationsCore{
					ValidatedAccount: true,
				},
			},
			expect: false,
		},
		{
			name:   "Success/NotFound",
			userID: goframework.NumberUUID(1),
			scope:  string(models.CanUseOpenAIPlayground),
			daoErr: bunovel.ErrNotFound,
			expect: false,
		},
		{
			name:      "Error/DAOFailure",
			userID:    goframework.NumberUUID(1),
			scope:     string(models.CanUseOpenAIPlayground),
			daoErr:    fooErr,
			expectErr: fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewUserAuthorizationsRepository(t)

			repository.
				On("Get", context.Background(), d.userID).
				Return(d.daoResp, d.daoErr)

			service := services.NewHasUserScopeService(repository)
			res, err := service.HasUserScope(context.Background(), d.userID, d.scope)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, res)

			repository.AssertExpectations(t)
		})
	}
}
