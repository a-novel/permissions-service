package handlers

import (
	"github.com/a-novel/authorizations-service/pkg/models"
	"github.com/a-novel/authorizations-service/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
)

type HasUserScopeHandler interface {
	Handle(c *gin.Context)
}

func NewHasUserScopeHandler(service services.HasUserScopeService) HasUserScopeHandler {
	return &hasUserScopeHandlerImpl{
		service: service,
	}
}

type hasUserScopeHandlerImpl struct {
	service services.HasUserScopeService
}

func (h *hasUserScopeHandlerImpl) Handle(c *gin.Context) {
	query := new(models.HasUserScopeQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	has, err := h.service.HasUserScope(c, query.UserID.Value(), query.Scope)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(lo.Ternary(has, http.StatusNoContent, http.StatusUnauthorized))
}
