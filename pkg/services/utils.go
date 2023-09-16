package services

import "errors"

var (
	ErrInvalidToken = errors.New("(data) invalid tokenRaw")

	ErrIntrospectToken = errors.New("(dep) failed to introspect token")

	ErrSetUserAuthorizations = errors.New("(dao) failed to set user authorizations")
	ErrGetUserAuthorizations = errors.New("(dao) failed to get user authorizations")
)
