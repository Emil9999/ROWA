package api

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

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

	if ArduinoOn {
		go DeactivateModuleLight()
	}
	return
}

func Harvest(c echo.Context) (err error) {
	fmt.Println("Harvest request received")
	plantType := new(PlantType)
	//Binding the received data to plantType
	c.Bind(plantType)

	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")

	sqlQuery := fmt.Sprintf("SELECT PlantPosition, Module FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and date(PlantDate, '+'||GrowthTime||' days') <= date('now') AND PlantType = '%s'", plantType.Name)

	rows, _ := database.Query(sqlQuery)
	//We only need the position (?)
	var position int
	var module int
	var arr []int
	for rows.Next() {
		rows.Scan(&position, &module)

		arr = append(arr, position)
		arr = append(arr, module)
	}
	arr = arr[:2]
	fmt.Println(arr)
	if ArduinoOn {
		go ActivateModuleLight(arr[1])
	}
	ret := c.JSON(http.StatusOK, arr)
	return ret

}
