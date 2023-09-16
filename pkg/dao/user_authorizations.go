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

type UserAuthorizationsRepository interface {
	Get(ctx context.Context, userID uuid.UUID) (*UserAuthorizations, error)
	Set(ctx context.Context, userID uuid.UUID, core UserAuthorizationsCore, fields AuthorizationFields, now time.Time) (*UserAuthorizations, error)
}

type UserAuthorizations struct {
	bun.BaseModel `bun:"table:users_authorizations"`
	bunovel.Metadata

	UserAuthorizationsCore
}

type UserAuthorizationsCore struct {
	ValidatedAccount bool `bun:"validated_account"`
	AdminAccess      bool `bun:"admin_access"`
}

type userAuthorizationsRepositoryImpl struct {
	db bun.IDB
}

func NewUserAuthorizationsRepository(db bun.IDB) UserAuthorizationsRepository {
	return &userAuthorizationsRepositoryImpl{
		db: db,
	}
}

func (repository *userAuthorizationsRepositoryImpl) Get(ctx context.Context, userID uuid.UUID) (*UserAuthorizations, error) {
	result := new(UserAuthorizations)

	err := repository.db.NewSelect().Model(result).Where("id = ?", userID).Scan(ctx)
	if err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return result, nil
}

func (repository *userAuthorizationsRepositoryImpl) Set(ctx context.Context, userID uuid.UUID, core UserAuthorizationsCore, fields AuthorizationFields, now time.Time) (*UserAuthorizations, error) {
	result := &UserAuthorizations{
		Metadata:               bunovel.NewMetadata(userID, now, nil),
		UserAuthorizationsCore: core,
	}

	query := repository.db.NewInsert().Model(result).Returning("*").
		On("CONFLICT (id) DO UPDATE").
		Set("updated_at = EXCLUDED.created_at")

	lo.ForEach(fields, func(item AuthorizationField, _ int) {
		query = query.Set(fmt.Sprintf("%[1]s = EXCLUDED.%[1]s", string(item)))
	})

	if err := query.Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return result, nil
}
