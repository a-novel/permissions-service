package adapters

import (
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/a-novel/permissions-service/pkg/models"
	"github.com/samber/lo"
)

var (
	DefaultScopes          = models.Scopes{}
	ValidatedAccountScopes = models.Scopes{
		models.CanVotePost,
		models.CanPostImproveRequest,
		models.CanPostImproveSuggestion,
	}
	AdminAccessScopes = models.Scopes{
		models.CanUseOpenAIPlayground,
	}

	PermissionsFields = dao.PermissionsFields{
		dao.FieldValidatedAccount,
		dao.FieldAdminAccess,
	}
)

func UserPermissionsModelToScopes(src *dao.UserPermissions) models.Scopes {
	scopes := models.Scopes{}
	scopes = append(scopes, DefaultScopes...)

	if src.ValidatedAccount {
		scopes = append(scopes, ValidatedAccountScopes...)
	}
	if src.AdminAccess {
		scopes = append(scopes, AdminAccessScopes...)
	}

	return scopes
}

func PermissionsFieldsToDAO(src []string) (dao.PermissionsFields, error) {
	var output dao.PermissionsFields

	for _, item := range src {
		toField := dao.PermissionField(item)
		if !lo.Contains(PermissionsFields, toField) {
			return nil, ErrInvalidPermissionField
		}

		output = append(output, toField)
	}

	return output, nil
}
