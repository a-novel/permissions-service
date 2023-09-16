package dao_test

import (
	"context"
	"github.com/a-novel/authorizations-service/migrations"
	"github.com/a-novel/authorizations-service/pkg/dao"
	"github.com/a-novel/bunovel"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"io/fs"
	"testing"
	"time"
)

func TestUsersAuthorizationsRepository_Get(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.UserAuthorizations{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UserAuthorizationsCore: dao.UserAuthorizationsCore{
				ValidatedAccount: true,
				AdminAccess:      false,
			},
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    *dao.UserAuthorizations
		expectErr error
	}{
		{
			name:   "Success",
			id:     goframework.NumberUUID(10),
			expect: fixtures[0].(*dao.UserAuthorizations),
		},
		{
			name:      "Error/Notfound",
			id:        goframework.NumberUUID(100),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewUserAuthorizationsRepository(tx)

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

func TestUsersAuthorizationsRepository_Set(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.UserAuthorizations{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
			UserAuthorizationsCore: dao.UserAuthorizationsCore{
				ValidatedAccount: true,
				AdminAccess:      false,
			},
		},
	}

	data := []struct {
		name string

		id     uuid.UUID
		core   dao.UserAuthorizationsCore
		fields dao.AuthorizationFields
		now    time.Time

		expect    *dao.UserAuthorizations
		expectErr error
	}{
		{
			name: "Success/Create",
			id:   goframework.NumberUUID(11),
			core: dao.UserAuthorizationsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.AuthorizationFields{dao.FieldAdminAccess},
			now:    updateTime,
			expect: &dao.UserAuthorizations{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(11), updateTime, nil),
				UserAuthorizationsCore: dao.UserAuthorizationsCore{
					ValidatedAccount: false,
					AdminAccess:      true,
				},
			},
		},
		{
			name: "Success/UpdateToTrue",
			id:   goframework.NumberUUID(10),
			core: dao.UserAuthorizationsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.AuthorizationFields{dao.FieldAdminAccess},
			now:    updateTime,
			expect: &dao.UserAuthorizations{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
				UserAuthorizationsCore: dao.UserAuthorizationsCore{
					ValidatedAccount: true,
					AdminAccess:      true,
				},
			},
		},
		{
			name: "Success/UpdateToFalse",
			id:   goframework.NumberUUID(10),
			core: dao.UserAuthorizationsCore{
				ValidatedAccount: false,
				AdminAccess:      true,
			},
			fields: dao.AuthorizationFields{dao.FieldValidatedAccount},
			now:    updateTime,
			expect: &dao.UserAuthorizations{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
				UserAuthorizationsCore: dao.UserAuthorizationsCore{
					ValidatedAccount: false,
					AdminAccess:      false,
				},
			},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(st *testing.T) {
			err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
				repository := dao.NewUserAuthorizationsRepository(tx)

				res, err := repository.Set(ctx, d.id, d.core, d.fields, d.now)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
			require.NoError(t, err)
		})
	}
}
