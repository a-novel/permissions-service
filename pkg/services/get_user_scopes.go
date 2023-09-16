package services

import (
	"context"
	"errors"
	"github.com/a-novel/bunovel"
	apiclients "github.com/a-novel/go-api-clients"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/adapters"
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/a-novel/permissions-service/pkg/models"
)

type GetUserScopesService interface {
	GetUserScopes(ctx context.Context, tokenRaw string) (models.Scopes, error)
}

func NewGetUserScopesService(repository dao.UserPermissionsRepository, authClient apiclients.AuthClient) GetUserScopesService {
	return &getUserScopesServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type getUserScopesServiceImpl struct {
	repository dao.UserPermissionsRepository
	authClient apiclients.AuthClient
}

func (s *getUserScopesServiceImpl) GetUserScopes(ctx context.Context, tokenRaw string) (models.Scopes, error) {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return nil, errors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return nil, errors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	permissions, err := s.repository.Get(ctx, token.Token.Payload.ID)
	if err != nil {
		if errors.Is(err, bunovel.ErrNotFound) {
			return nil, nil
		}

		return nil, errors.Join(ErrGetUserPermissions, err)
	}

	return adapters.UserPermissionsModelToScopes(permissions), nil
}
