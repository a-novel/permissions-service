package services_test

import (
	"context"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/dao"
	daomocks "github.com/a-novel/permissions-service/pkg/dao/mocks"
	"github.com/a-novel/permissions-service/pkg/models"
	"github.com/a-novel/permissions-service/pkg/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSetUserPermissionsService(t *testing.T) {
	data := []struct {
		name string

		userID      uuid.UUID
		setFields   []string
		unsetFields []string
		now         time.Time

		shouldCallDAO           bool
		shouldCallDAOWithCore   dao.UserPermissionsCore
		shouldCallDAOWithFields dao.PermissionsFields
		daoErr                  error

		expectErr error
	}{
		{
			name:   "Success",
			userID: goframework.NumberUUID(1),
			setFields: []string{
				string(models.FieldValidatedAccount),
			},
			unsetFields: []string{
				string(models.FieldAdminAccess),
			},
			now:           baseTime,
			shouldCallDAO: true,
			shouldCallDAOWithCore: dao.UserPermissionsCore{
				ValidatedAccount: true,
			},
			shouldCallDAOWithFields: dao.PermissionsFields{
				dao.FieldValidatedAccount,
				dao.FieldAdminAccess,
			},
		},
		{
			name:   "Success/NoFields",
			userID: goframework.NumberUUID(1),
			now:    baseTime,
		},
		{
			name:   "Error/InvalidSetField",
			userID: goframework.NumberUUID(1),
			setFields: []string{
				string(models.FieldValidatedAccount),
				"invalid",
			},
			unsetFields: []string{
				string(models.FieldAdminAccess),
			},
			now:       baseTime,
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:   "Error/InvalidUnsetField",
			userID: goframework.NumberUUID(1),
			setFields: []string{
				string(models.FieldValidatedAccount),
			},
			unsetFields: []string{
				string(models.FieldAdminAccess),
				"invalid",
			},
			now:       baseTime,
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:   "Error/DAOFailure",
			userID: goframework.NumberUUID(1),
			setFields: []string{
				string(models.FieldValidatedAccount),
			},
			unsetFields: []string{
				string(models.FieldAdminAccess),
			},
			now:           baseTime,
			shouldCallDAO: true,
			shouldCallDAOWithCore: dao.UserPermissionsCore{
				ValidatedAccount: true,
			},
			shouldCallDAOWithFields: dao.PermissionsFields{
				dao.FieldValidatedAccount,
				dao.FieldAdminAccess,
			},
			daoErr:    fooErr,
			expectErr: fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewUserPermissionsRepository(t)

			if d.shouldCallDAO {
				repository.
					On("Set", context.Background(), d.userID, d.shouldCallDAOWithCore, d.shouldCallDAOWithFields, d.now).
					Return(nil, d.daoErr)
			}

			service := services.NewSetUserPermissionsService(repository)
			err := service.Set(context.Background(), d.userID, d.setFields, d.unsetFields, d.now)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
		})
	}
}
