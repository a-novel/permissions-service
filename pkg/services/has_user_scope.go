package services

import (
	"context"
	"errors"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/permissions-service/pkg/adapters"
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/a-novel/permissions-service/pkg/models"
	"github.com/google/uuid"
)

type HasUserScopeService interface {
	HasUserScope(ctx context.Context, userID uuid.UUID, scope string) (bool, error)
}

func NewHasUserScopeService(repository dao.UserPermissionsRepository) HasUserScopeService {
	return &hasUserScopeServiceImpl{
		repository: repository,
	}
}

type hasUserScopeServiceImpl struct {
	repository dao.UserPermissionsRepository
}

func (s *hasUserScopeServiceImpl) HasUserScope(ctx context.Context, userID uuid.UUID, scope string) (bool, error) {
	permissions, err := s.repository.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, bunovel.ErrNotFound) {
			return false, nil
		}

		return false, errors.Join(ErrGetUserPermissions, err)
	}

	scopes := adapters.UserPermissionsModelToScopes(permissions)

	return scopes.Has(models.Scope(scope)), nil
}
