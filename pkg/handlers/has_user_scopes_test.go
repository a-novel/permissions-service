package handlers_test

import (
	goframework "github.com/a-novel/go-framework"
	"github.com/a-novel/permissions-service/pkg/handlers"
	servicesmocks "github.com/a-novel/permissions-service/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHasUserScopesHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService           bool
		shouldCallServiceWithScope  string
		shouldCallServiceWithUserID uuid.UUID
		serviceResp                 bool
		serviceErr                  error

		expectStatus int
	}{
		{
			name:                        "Success",
			query:                       "?scope=foo&userID=01010101-0101-0101-0101-010101010101",
			shouldCallServiceWithScope:  "foo",
			shouldCallServiceWithUserID: goframework.NumberUUID(1),
			shouldCallService:           true,
			serviceResp:                 true,
			expectStatus:                http.StatusNoContent,
		},
		{
			name:                        "Success/NotAuthorized",
			query:                       "?scope=foo&userID=01010101-0101-0101-0101-010101010101",
			shouldCallServiceWithScope:  "foo",
			shouldCallServiceWithUserID: goframework.NumberUUID(1),
			shouldCallService:           true,
			serviceResp:                 false,
			expectStatus:                http.StatusUnauthorized,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewHasUserScopeService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			service.
				On("HasUserScope", c, d.shouldCallServiceWithUserID, d.shouldCallServiceWithScope).
				Return(d.serviceResp, d.serviceErr)

			handler := handlers.NewHasUserScopeHandler(service)
			handler.Handle(c)

			require.Equal(t, d.expectStatus, w.Code, c.Errors.String())

			service.AssertExpectations(t)
		})
	}
}
