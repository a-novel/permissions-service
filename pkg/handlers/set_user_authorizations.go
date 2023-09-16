package handlers

import (
	"github.com/a-novel/authorizations-service/pkg/models"
	"github.com/a-novel/authorizations-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SetUserAuthorizationsHandler interface {
	Handle(c *gin.Context)
}

func NewSetUserAuthorizationsHandler(service services.SetUserAuthorizationService) SetUserAuthorizationsHandler {
	return &setUserAuthorizationsHandlerImpl{
		service: service,
	}
}

type setUserAuthorizationsHandlerImpl struct {
	service services.SetUserAuthorizationService
}

func (h *setUserAuthorizationsHandlerImpl) Handle(c *gin.Context) {
	form := new(models.SetUserAuthorizationsForm)
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
