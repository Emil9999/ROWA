package api

import (
	"db"
	"github.com/labstack/echo"
	"net/http"
)

func GetHarvestablePlantsHandler(c echo.Context) (err error) {
	plantsToHarvest, err := db.FunctionStore.GetPlantsPerType("harvestable")
	if err != nil {
		return c.JSON(http.StatusNotFound, "Harvestable Plants not found")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func GetPlantablePlantsHandler(c echo.Context) (err error) {
	plantsToHarvest, err := db.FunctionStore.GetPlantsPerType("plantable")
	if err != nil {
		return c.JSON(http.StatusNotFound, "Plantable Plants not found")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func GetSensorDataHandler(c echo.Context) (err error) {
	sensorData, err := db.FunctionStore.GetLastSensorEntry()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Sensor Data not found")
	}

	return c.JSON(http.StatusOK, sensorData)
}

// TODO Function for CatTree Information Handler @Emil, @Behnaz
