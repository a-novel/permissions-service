package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/a-novel/authorizations-service/pkg/handlers"
	servicesmocks "github.com/a-novel/authorizations-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetUserAuthorizationsHandler(t *testing.T) {
	data := []struct {
		name string

		body interface{}

		shouldCallService                bool
		shouldCallServiceWithUserID      uuid.UUID
		shouldCallServiceWithSetFields   []string
		shouldCallServiceWithUnsetFields []string
		serviceErr                       error

		expectStatus int
	}{
		{
			name: "Success",
			body: map[string]interface{}{
				"userID":      goframework.NumberUUID(1),
				"setFields":   []string{"foo", "bar"},
				"unsetFields": []string{"baz", "qux"},
			},
			shouldCallService:                true,
			shouldCallServiceWithUserID:      goframework.NumberUUID(1),
			shouldCallServiceWithSetFields:   []string{"foo", "bar"},
			shouldCallServiceWithUnsetFields: []string{"baz", "qux"},
			expectStatus:                     http.StatusCreated,
		},
		{
			name: "Error/InvalidEntity",
			body: map[string]interface{}{
				"userID":      goframework.NumberUUID(1),
				"setFields":   []string{"foo", "bar"},
				"unsetFields": []string{"baz", "qux"},
			},
			shouldCallService:                true,
			shouldCallServiceWithUserID:      goframework.NumberUUID(1),
			shouldCallServiceWithSetFields:   []string{"foo", "bar"},
			shouldCallServiceWithUnsetFields: []string{"baz", "qux"},
			serviceErr:                       goframework.ErrInvalidEntity,
			expectStatus:                     http.StatusUnprocessableEntity,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewSetUserAuthorizationService(t)

			mrshBody, err := json.Marshal(d.body)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(mrshBody))

			if d.shouldCallService {
				service.
					On("Set", c, d.shouldCallServiceWithUserID, d.shouldCallServiceWithSetFields, d.shouldCallServiceWithUnsetFields, mock.Anything).
					Return(d.serviceErr)
			}

			handler := handlers.NewSetUserAuthorizationsHandler(service)
			handler.Handle(c)

			require.Equal(t, d.expectStatus, w.Code, c.Errors.String())

			service.AssertExpectations(t)
		})
	}
}
