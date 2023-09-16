package services

import (
	"context"
	"errors"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/adapters"
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"time"
)

type SetUserPermissionsService interface {
	Set(ctx context.Context, userID uuid.UUID, setFields, unsetFields []string, now time.Time) error
}

func NewSetUserPermissionsService(repository dao.UserPermissionsRepository) SetUserPermissionsService {
	return &setUserPermissionsServiceImpl{
		repository: repository,
	}
}

type setUserPermissionsServiceImpl struct {
	repository dao.UserPermissionsRepository
}

func (s *setUserPermissionsServiceImpl) Set(ctx context.Context, userID uuid.UUID, setFields, unsetFields []string, now time.Time) error {
	setFieldsDAO, err := adapters.PermissionsFieldsToDAO(setFields)
	if err != nil {
		return errors.Join(goframework.ErrInvalidEntity, err)
	}
	unsetFieldsDAO, err := adapters.PermissionsFieldsToDAO(unsetFields)
	if err != nil {
		return errors.Join(goframework.ErrInvalidEntity, err)
	}

	if len(setFieldsDAO) == 0 && len(unsetFieldsDAO) == 0 {
		return nil
	}

	core := dao.UserPermissionsCore{}

	lo.ForEach(setFieldsDAO, func(item dao.PermissionField, _ int) {
		switch item {
		case dao.FieldValidatedAccount:
			core.ValidatedAccount = true
		case dao.FieldAdminAccess:
			core.AdminAccess = true
		}
	})

	_, err = s.repository.Set(ctx, userID, core, append(setFieldsDAO, unsetFieldsDAO...), now)
	if err != nil {
		return errors.Join(ErrSetUserPermissions, err)
	}

	return nil
}
