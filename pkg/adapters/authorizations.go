package adapters

import (
	"github.com/a-novel/authorizations-service/pkg/dao"
	"github.com/a-novel/authorizations-service/pkg/models"
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

	AuthorizationsFields = dao.AuthorizationFields{
		dao.FieldValidatedAccount,
		dao.FieldAdminAccess,
	}
)

func UserAuthorizationsModelToScopes(src *dao.UserAuthorizations) models.Scopes {
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

func AuthorizationsFieldsToDAO(src []string) (dao.AuthorizationFields, error) {
	var output dao.AuthorizationFields

	for _, item := range src {
		toField := dao.AuthorizationField(item)
		if !lo.Contains(AuthorizationsFields, toField) {
			return nil, ErrInvalidAuthorizationField
		}

		output = append(output, toField)
	}

	return output, nil
}
