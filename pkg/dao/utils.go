package dao

type AuthorizationField string

type AuthorizationFields []AuthorizationField

const (
	FieldValidatedAccount AuthorizationField = "validated_account"
	FieldAdminAccess      AuthorizationField = "admin_access"
)
