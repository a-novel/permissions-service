package models

import "github.com/google/uuid"

type SetUserPermissionsForm struct {
	UserID      uuid.UUID `json:"userID" form:"userID"`
	SetFields   []string  `json:"setFields" form:"setFields"`
	UnsetFields []string  `json:"unsetFields" form:"unsetFields"`
}
