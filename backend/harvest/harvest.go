package harvest

import (
	"database/sql"
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

func HarvestDone(c echo.Context) (err error) {
	fmt.Println("Harvested!")
	plantPosition := new(PositionOnFarm)
	c.Bind(plantPosition)
	fmt.Println(plantPosition)
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")

	sqlQuery := fmt.Sprintf("UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = '%d' AND Module=%d", plantPosition.PlantPosition, plantPosition.ModulePosition)
	statement, _ := database.Prepare(sqlQuery)
	statement.Exec()
	sqlQuery = fmt.Sprintf("UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position='%d'", plantPosition.ModulePosition)
	statement, _ = database.Prepare(sqlQuery)
	statement.Exec()

	if settings.ArduinoOn {
		go sensors.DeactivateModuleLight()
	}
	return
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
