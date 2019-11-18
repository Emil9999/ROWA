package harvest

import (
	"fmt"
	"net/http"

	"../sensors"
	"../settings"
	"../utils"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

type DB utils.DbObject

type PositionOnFarm struct {
	PlantPosition  int `json: "plantPosition"`
	ModulePosition int `json: "modulePosition"`
}

func HarvestDone(c echo.Context, db *utils.DbObject) (err error) {
	fmt.Println("Harvested!")
	plantPosition := new(PositionOnFarm)
	c.Bind(plantPosition)
	fmt.Println(plantPosition)
	success := utils.UpdateAfterHarvest(plantPosition.PlantPosition, plantPosition.ModulePosition, db)
	if settings.ArduinoOn && success {
		go sensors.DeactivateModuleLight()
	}
	return err
}

func Harvest(c echo.Context, db *utils.DbObject) (err error) {
	fmt.Println("Harvest request received")
	plantType := new(utils.PlantType)
	//Binding the received data to plantType
	c.Bind(plantType)
	//Opening DB
	arr := utils.GetHarvestablePlantQuery(plantType.Name, db)
	fmt.Println(arr)
	if settings.ArduinoOn {
		go sensors.ActivateModuleLight(arr[1])
	}
	ret := c.JSON(http.StatusOK, arr)
	return ret

}
