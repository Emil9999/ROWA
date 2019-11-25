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

func TestPlantHandler(t *testing.T) {
	mockStore := db.InitMockStore()

	mockStore.On("Plant", &db.PlantType{}).Return(1, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/plant/get-position", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := 1

	if assert.NoError(t, PlantHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody int
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}

func TestFinishPlantingHandler(t *testing.T) {
	mockStore := db.InitMockStore()

	mockStore.On("FinishPlanting", &db.PlantedModule{}).Return(
		&db.Status{Message: "Planting Done"}, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/plant/finish", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.Status{Message: "Planting Done"}

	if assert.NoError(t, FinishPlantingHandler(c)) {
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
