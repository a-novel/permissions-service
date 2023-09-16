package models

import "github.com/a-novel/go-apis"

type HasUserScopeQuery struct {
	UserID apis.StringUUID `json:"userID" form:"userID"`
	Scope  string          `json:"scope" form:"scope"`
}
