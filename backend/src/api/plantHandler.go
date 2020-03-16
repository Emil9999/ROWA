package api

import (
	"net/http"
    "github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/db"
	"github.com/labstack/echo"
)

func PlantHandler(c echo.Context) (err error) {
	plantType := &db.PlantType{}
	err = c.Bind(plantType)
	position, err := db.FunctionStore.Plant(plantType)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Planting not available")
	}

	return c.JSON(http.StatusOK, position)
}

func FinishPlantingHandler(c echo.Context) (err error) {
	plantedModule := &db.PlantedModule{}
	c.Bind(&plantedModule)

	status, err := db.FunctionStore.FinishPlanting(plantedModule)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, status)
}

func StopModuleBlink(c echo.Context) (err error) {

	sensor.DeactivateModuleLight()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, "Canceled Light")
}
