package handlers

import (
	"github.com/a-novel/authorizations-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserScopesHandler interface {
	Handle(c *gin.Context)
}

func NewGetUserScopesHandler(service services.GetUserScopesService) GetUserScopesHandler {
	return &getUserScopesHandlerImpl{
		service: service,
	}
}

type getUserScopesHandlerImpl struct {
	service services.GetUserScopesService
}

func (h *getUserScopesHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	scopes, err := h.service.GetUserScopes(c, token)
	if err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{goframework.ErrInvalidCredentials, http.StatusForbidden},
		}, false)
		return
	}

	c.JSON(200, gin.H{"scopes": scopes})
}
