package health_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielteiga/startup-rush/api/controllers/health"
	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/stretchr/testify/assert"
)

const HEALTH_ENDPOINT = "/api/health"

func TestHealthWorks(t *testing.T) {
	var resp *responses.Response[any]

	req := httptest.NewRequest(http.MethodGet, HEALTH_ENDPOINT, nil)
	rec := httptest.NewRecorder()

	health.NewHealthController().Handle(rec, req)

	err := json.NewDecoder(rec.Body).Decode(&resp)
	assert.NoError(t, err)

	expected := &responses.Response[any]{
		Message: "The app is healthy!",
		Status:  "Success",
	}

	assert.Equal(t, expected, resp)
}
