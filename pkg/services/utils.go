package services

import "errors"

var (
	ErrInvalidToken = errors.New("(data) invalid tokenRaw")

	ErrIntrospectToken = errors.New("(dep) failed to introspect token")

	ErrSetUserPermissions = errors.New("(dao) failed to set user permissions")
	ErrGetUserPermissions = errors.New("(dao) failed to get user permissions")
)
