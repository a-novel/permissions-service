package dao_test

import (
	"context"
	"github.com/a-novel/bunovel"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/migrations"
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"io/fs"
	"testing"
	"time"
)

func TestUsersPermissionsRepository_Get(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.UserPermissions{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UserPermissionsCore: dao.UserPermissionsCore{
				ValidatedAccount: true,
				AdminAccess:      false,
			},
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    *dao.UserPermissions
		expectErr error
	}{
		{
			name:   "Success",
			id:     goframework.NumberUUID(10),
			expect: fixtures[0].(*dao.UserPermissions),
		},
		{
			name:      "Error/Notfound",
			id:        goframework.NumberUUID(100),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewUserPermissionsRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Get(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
		}
	})
	require.NoError(t, err)
}

func TestUsersPermissionsRepository_Set(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.UserPermissions{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
			UserPermissionsCore: dao.UserPermissionsCore{
				ValidatedAccount: true,
				AdminAccess:      false,
			},
		},
	}

	data := []struct {
		name string

		id     uuid.UUID
		core   dao.UserPermissionsCore
		fields dao.PermissionsFields
		now    time.Time

		expect    *dao.UserPermissions
		expectErr error
	}{
		{
			name: "Success/Create",
			id:   goframework.NumberUUID(11),
			core: dao.UserPermissionsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.PermissionsFields{dao.FieldAdminAccess},
			now:    updateTime,
			expect: &dao.UserPermissions{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(11), updateTime, nil),
				UserPermissionsCore: dao.UserPermissionsCore{
					ValidatedAccount: false,
					AdminAccess:      true,
				},
			},
		},
		{
			name: "Success/UpdateToTrue",
			id:   goframework.NumberUUID(10),
			core: dao.UserPermissionsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.PermissionsFields{dao.FieldAdminAccess},
			now:    updateTime,
			expect: &dao.UserPermissions{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
				UserPermissionsCore: dao.UserPermissionsCore{
					ValidatedAccount: true,
					AdminAccess:      true,
				},
			},
		},
		{
			name: "Success/UpdateToFalse",
			id:   goframework.NumberUUID(10),
			core: dao.UserPermissionsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.PermissionsFields{dao.FieldValidatedAccount},
			now:    updateTime,
			expect: &dao.UserPermissions{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
				UserPermissionsCore: dao.UserPermissionsCore{
					ValidatedAccount: false,
					AdminAccess:      false,
				},
			},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(st *testing.T) {
			err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
				repository := dao.NewUserPermissionsRepository(tx)

				res, err := repository.Set(ctx, d.id, d.core, d.fields, d.now)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
			require.NoError(t, err)
		})
	}
}
