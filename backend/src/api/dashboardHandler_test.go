package api

import (
	"db"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSensorDataHandler(t *testing.T) {
	mockStore := db.InitMockStore()

	mockStore.On("GetLastSensorEntry").Return(
		&db.SensorData{Datetime: "2018-06-05", Temp: 2.4, LightIntensity: 120}, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/dashboard/sensor-data", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.SensorData{Datetime: "2018-06-05", Temp: 2.4, LightIntensity: 120}

	if assert.NoError(t, GetSensorDataHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		requestBody := &db.SensorData{}
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}
