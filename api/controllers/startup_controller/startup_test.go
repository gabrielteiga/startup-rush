package startup_controller_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielteiga/startup-rush/api/controllers/startup_controller"
	"github.com/gabrielteiga/startup-rush/api/requests"
	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
	"github.com/gabrielteiga/startup-rush/internal/infra/repositories/gorm_adapter"
	"github.com/gabrielteiga/startup-rush/internal/utils/parsers"
	"github.com/stretchr/testify/assert"
)

const CREATE_STARTUP_ENDPOINT string = "/api/startups"

var dbConn = database.InitConnection()

var StartupRepository startup_entity.IStartupRepository = gorm_adapter.NewStartupGORMRepository(dbConn.DB)
var StartupService *services.StartupService = services.NewStartupService(StartupRepository)

func TestCreateStartupWorks(t *testing.T) {
	DPStartupTest := []struct {
		Request        *requests.RequestStartupCreate
		ExpectedResult string
	}{
		{
			Request:        requests.NewRequestStartupCreate("Dell Technologies Brasil", "Always forward!", "2021-10-20T15:04:05Z"),
			ExpectedResult: responses.SUCCESS_STATUS,
		},
		{
			Request:        requests.NewRequestStartupCreate("Dell Technologies USA", "Keep trying!", "2021-10-20T15:04:05Z"),
			ExpectedResult: responses.SUCCESS_STATUS,
		},
		{
			Request:        requests.NewRequestStartupCreate("Dell Technologies India", "Do your best!", "2021-10-20T15:04:05Z"),
			ExpectedResult: responses.SUCCESS_STATUS,
		},
	}

	for i, data := range DPStartupTest {
		parsedFoundation, err := parsers.StringDateToTime(data.Request.Foundation)
		if err != nil {
			log.Println(err)
		}

		startup := &startup_entity.Startup{
			ID:         uint(i),
			Name:       data.Request.Name,
			Slogan:     data.Request.Slogan,
			Foundation: parsedFoundation,
		}

		var expectedResult *responses.Response[startup_entity.Startup]

		switch data.ExpectedResult {
		case responses.SUCCESS_STATUS:
			expectedResult = responses.NewReponse(
				startup_controller.CREATED_SUCCESSFULLY_MESSAGE,
				responses.SUCCESS_STATUS,
				startup,
			)
		default:
			expectedResult = responses.NewReponse(
				startup_controller.CREATED_ERROR_MESSAGE,
				responses.ERROR_STATUS,
				startup,
			)
			expectedResult.Data = nil
		}

		createStartupWorks(t, data.Request, expectedResult)
	}
}

func createStartupWorks(t *testing.T, request *requests.RequestStartupCreate, expectedResult *responses.Response[startup_entity.Startup]) {
	var resp *responses.Response[startup_entity.Startup]
	bodyRequest, _ := json.Marshal(request)

	req := httptest.NewRequest(http.MethodPost, CREATE_STARTUP_ENDPOINT, bytes.NewReader(bodyRequest))
	rec := httptest.NewRecorder()

	startup_controller.NewStartupController(StartupService).Handle(rec, req)

	err := json.NewDecoder(rec.Body).Decode(&resp)
	assert.NoError(t, err)

	if expectedResult.Status == responses.SUCCESS_STATUS {
		resp.Data.ID = expectedResult.Data.ID
	}

	assert.Equal(t, expectedResult, resp)
}
