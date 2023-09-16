package services

import (
	"context"
	"errors"
	"github.com/a-novel/authorizations-service/pkg/adapters"
	"github.com/a-novel/authorizations-service/pkg/dao"
	"github.com/a-novel/authorizations-service/pkg/models"
	"github.com/a-novel/bunovel"
	apiclients "github.com/a-novel/go-api-clients"
	goframework "github.com/a-novel/go-framework"
)

type GetUserScopesService interface {
	GetUserScopes(ctx context.Context, tokenRaw string) (models.Scopes, error)
}

func NewGetUserScopesService(repository dao.UserAuthorizationsRepository, authClient apiclients.AuthClient) GetUserScopesService {
	return &getUserScopesServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type getUserScopesServiceImpl struct {
	repository dao.UserAuthorizationsRepository
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

	authorizations, err := s.repository.Get(ctx, token.Token.Payload.ID)
	if err != nil {
		if errors.Is(err, bunovel.ErrNotFound) {
			return nil, nil
		}

		return nil, errors.Join(ErrGetUserAuthorizations, err)
	}

	return adapters.UserAuthorizationsModelToScopes(authorizations), nil
}
