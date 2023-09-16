package services

import (
	"context"
	"errors"
	"github.com/a-novel/authorizations-service/pkg/adapters"
	"github.com/a-novel/authorizations-service/pkg/dao"
	"github.com/a-novel/authorizations-service/pkg/models"
	"github.com/a-novel/bunovel"
	"github.com/google/uuid"
)

type HasUserScopeService interface {
	HasUserScope(ctx context.Context, userID uuid.UUID, scope string) (bool, error)
}

func NewHasUserScopeService(repository dao.UserAuthorizationsRepository) HasUserScopeService {
	return &hasUserScopeServiceImpl{
		repository: repository,
	}
}

type hasUserScopeServiceImpl struct {
	repository dao.UserAuthorizationsRepository
}

func (s *hasUserScopeServiceImpl) HasUserScope(ctx context.Context, userID uuid.UUID, scope string) (bool, error) {
	authorizations, err := s.repository.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, bunovel.ErrNotFound) {
			return false, nil
		}

		return false, errors.Join(ErrGetUserAuthorizations, err)
	}

	scopes := adapters.UserAuthorizationsModelToScopes(authorizations)

	return scopes.Has(models.Scope(scope)), nil
}
