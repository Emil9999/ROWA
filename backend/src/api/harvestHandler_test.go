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

func TestGetHarvestablePlantHandler(t *testing.T) {
	mockStore := db.InitMockStore()

	mockStore.On("GetHarvestablePlant", &db.PlantType{}).Return(
		&db.PositionOnFarm{PlantPosition: 1, ModulePosition: 1}, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/harvest/get-plant", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.PositionOnFarm{PlantPosition: 1, ModulePosition: 1}

	if assert.NoError(t, GetHarvestablePlantHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.PositionOnFarm
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}

func TestHarvestDoneHandler(t *testing.T) {
	mockStore := db.InitMockStore()

	mockStore.On("HarvestDone", &db.PositionOnFarm{}).Return(
		&db.Status{Message: "Harvest Done"}, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/harvest/harvestdone", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.Status{Message: "Harvest Done"}

	if assert.NoError(t, HarvestDoneHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.Status
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}
