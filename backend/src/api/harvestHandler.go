package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/db"

	"github.com/labstack/echo"
)

func GetHarvestablePlantHandler(c echo.Context) (err error) {
	plantType := &db.PlantType{}
	//Binding the received data to plantType
	err = c.Bind(plantType)
	plantsToHarvest, err := db.FunctionStore.GetHarvestablePlant(plantType)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Plant type not available")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func HarvestDoneHandler(c echo.Context) (err error) {
	plantPosition := new(db.PositionOnFarm)

	err = c.Bind(plantPosition)
	fmt.Println(plantPosition)
	fmt.Println(err)
	status, err := db.FunctionStore.HarvestDone(plantPosition)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Harvest not possible")
	}

	return c.JSON(http.StatusOK, status)
}
