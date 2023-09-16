package services

import (
	"context"
	"errors"
	"github.com/a-novel/authorizations-service/pkg/adapters"
	"github.com/a-novel/authorizations-service/pkg/dao"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"time"
)

type SetUserAuthorizationService interface {
	Set(ctx context.Context, userID uuid.UUID, setFields, unsetFields []string, now time.Time) error
}

func NewSetUserAuthorizationService(repository dao.UserAuthorizationsRepository) SetUserAuthorizationService {
	return &setUserAuthorizationServiceImpl{
		repository: repository,
	}
}

type setUserAuthorizationServiceImpl struct {
	repository dao.UserAuthorizationsRepository
}

func (s *setUserAuthorizationServiceImpl) Set(ctx context.Context, userID uuid.UUID, setFields, unsetFields []string, now time.Time) error {
	setFieldsDAO, err := adapters.AuthorizationsFieldsToDAO(setFields)
	if err != nil {
		return errors.Join(goframework.ErrInvalidEntity, err)
	}
	unsetFieldsDAO, err := adapters.AuthorizationsFieldsToDAO(unsetFields)
	if err != nil {
		return errors.Join(goframework.ErrInvalidEntity, err)
	}

	if len(setFieldsDAO) == 0 && len(unsetFieldsDAO) == 0 {
		return nil
	}

	core := dao.UserAuthorizationsCore{}

	lo.ForEach(setFieldsDAO, func(item dao.AuthorizationField, _ int) {
		switch item {
		case dao.FieldValidatedAccount:
			core.ValidatedAccount = true
		case dao.FieldAdminAccess:
			core.AdminAccess = true
		}
	})

	_, err = s.repository.Set(ctx, userID, core, append(setFieldsDAO, unsetFieldsDAO...), now)
	if err != nil {
		return errors.Join(ErrSetUserAuthorizations, err)
	}

	return nil
}
