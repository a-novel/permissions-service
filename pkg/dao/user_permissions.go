package dao

import (
	"context"
	"fmt"
	"github.com/a-novel/bunovel"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
	"time"
)

type UserPermissionsRepository interface {
	Get(ctx context.Context, userID uuid.UUID) (*UserPermissions, error)
	Set(ctx context.Context, userID uuid.UUID, core UserPermissionsCore, fields PermissionsFields, now time.Time) (*UserPermissions, error)
}

type UserPermissions struct {
	bun.BaseModel `bun:"table:users_permissions"`
	bunovel.Metadata

	UserPermissionsCore
}

type UserPermissionsCore struct {
	ValidatedAccount bool `bun:"validated_account"`
	AdminAccess      bool `bun:"admin_access"`
}

type userPermissionsRepositoryImpl struct {
	db bun.IDB
}

func NewUserPermissionsRepository(db bun.IDB) UserPermissionsRepository {
	return &userPermissionsRepositoryImpl{
		db: db,
	}
}

func (repository *userPermissionsRepositoryImpl) Get(ctx context.Context, userID uuid.UUID) (*UserPermissions, error) {
	result := new(UserPermissions)

	err := repository.db.NewSelect().Model(result).Where("id = ?", userID).Scan(ctx)
	if err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return result, nil
}

func (repository *userPermissionsRepositoryImpl) Set(ctx context.Context, userID uuid.UUID, core UserPermissionsCore, fields PermissionsFields, now time.Time) (*UserPermissions, error) {
	result := &UserPermissions{
		Metadata:            bunovel.NewMetadata(userID, now, nil),
		UserPermissionsCore: core,
	}

	query := repository.db.NewInsert().Model(result).Returning("*").
		On("CONFLICT (id) DO UPDATE").
		Set("updated_at = EXCLUDED.created_at")

	lo.ForEach(fields, func(item PermissionField, _ int) {
		query = query.Set(fmt.Sprintf("%[1]s = EXCLUDED.%[1]s", string(item)))
	})

	if err := query.Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return result, nil
}
