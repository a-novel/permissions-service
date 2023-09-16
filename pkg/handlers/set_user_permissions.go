package handlers

import (
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/models"
	"github.com/a-novel/permissions-service/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SetUserPermissionsHandler interface {
	Handle(c *gin.Context)
}

func NewSetUserPermissionsHandler(service services.SetUserPermissionsService) SetUserPermissionsHandler {
	return &setUserPermissionsHandlerImpl{
		service: service,
	}
}

type setUserPermissionsHandlerImpl struct {
	service services.SetUserPermissionsService
}

func (h *setUserPermissionsHandlerImpl) Handle(c *gin.Context) {
	form := new(models.SetUserPermissionsForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.Set(c, form.UserID, form.SetFields, form.UnsetFields, time.Now()); err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{goframework.ErrInvalidEntity, http.StatusUnprocessableEntity},
		}, false)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
