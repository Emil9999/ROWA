package api

import (
	"db"
	"github.com/labstack/echo"
	"net/http"
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

	err = db.FunctionStore.FinishPlanting(plantedModule)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, true)
}
