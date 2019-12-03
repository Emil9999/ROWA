package api

import (
	"db"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	_ "golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitialiseTestServer(httpMethod string, url string) (c echo.Context, rec *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(httpMethod, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	return
}

func TestGetSensorDataHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("GetLastSensorEntry").Return(
		&db.SensorData{Datetime: "2018-06-05", Temp: 2.4, LightIntensity: 120}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/sensor-data")

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

func TestGetHarvestablePlantsHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("GetPlantsPerType", "harvestable").Return(
		[]*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/harvestable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}

	if assert.NoError(t, GetHarvestablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}

func TestGetPlantablePlantsHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("GetPlantsPerType", "plantable").Return(
		[]*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/plantable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}

	if assert.NoError(t, GetPlantablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}
