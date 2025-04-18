package health_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielteiga/startup-rush/api/controllers/health"
	"github.com/stretchr/testify/assert"
)

const HEALTH_ENDPOINT = "/api/health"

type ResponseHealth struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func TestHealthWorks(t *testing.T) {
	var resp *ResponseHealth

	req := httptest.NewRequest(http.MethodGet, HEALTH_ENDPOINT, nil)
	rec := httptest.NewRecorder()

	health.NewHealthController().Handle(rec, req)

	body, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	json.Unmarshal(body, &resp)

	expected := &ResponseHealth{
		Status:  "Success",
		Message: "The app is healthy!",
	}

	assert.Equal(t, expected, resp)
}
