package dao

type PermissionField string

type PermissionsFields []PermissionField

const (
	FieldValidatedAccount PermissionField = "validated_account"
	FieldAdminAccess      PermissionField = "admin_access"
)
