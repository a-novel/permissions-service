package handlers_test

import (
	"encoding/json"
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/handlers"
	"github.com/a-novel/permissions-service/pkg/models"
	servicesmocks "github.com/a-novel/permissions-service/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserScopesHandler(t *testing.T) {
	data := []struct {
		name string

		authorization string

		serviceResp models.Scopes
		serviceErr  error

		expect       interface{}
		expectStatus int
	}{
		{
			name:          "Success",
			authorization: "Bearer my-token",
			serviceResp:   models.Scopes{"foo", "bar"},
			expect:        map[string]interface{}{"scopes": []interface{}{"foo", "bar"}},
			expectStatus:  http.StatusOK,
		},
		{
			name:          "Error/ErrInvalidCredentials",
			authorization: "Bearer my-token",
			serviceErr:    goframework.ErrInvalidCredentials,
			expectStatus:  http.StatusForbidden,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewGetUserScopesService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", nil)
			c.Request.Header.Set("Authorization", d.authorization)

			service.
				On("GetUserScopes", c, d.authorization).
				Return(d.serviceResp, d.serviceErr)

			handler := handlers.NewGetUserScopesHandler(service)
			handler.Handle(c)

			require.Equal(t, d.expectStatus, w.Code, c.Errors.String())
			if d.expect != nil {
				var body interface{}
				require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
				require.Equal(t, d.expect, body)
			}

			service.AssertExpectations(t)
		})
	}
}
