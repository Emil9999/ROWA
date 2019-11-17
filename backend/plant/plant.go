package plant

import (
	"../sensors"
	"../settings"
	"../utils"
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type PlantedModule struct {
	Module int `json:"PlantedModule"`
}

func Plant(c echo.Context) (err error) {
	fmt.Println("This is server: request received")

	plantType := new(utils.PlantType)
	//Binding the received data to plantType
	c.Bind(plantType)

	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Selecting the modules that have available spots and match the plant type
	sqlQuery := fmt.Sprintf("SELECT Position FROM Module WHERE AvailableSpots>0 AND PlantType='%s'", plantType.Name)
	rows, _ := database.Query(sqlQuery)
	//We only need the position (?)
	var position int
	for rows.Next() {
		rows.Scan(&position)
		fmt.Println(strconv.Itoa(position))
	}
	// Light up module
	if settings.ArduinoOn{
		go sensors.ActivateModuleLight(position)
	}

	//Returning only the last position since user can only plant in one module
	ret := c.JSON(http.StatusOK, position)
	return ret
}

func FinishPlanting(c echo.Context) (err error) {

	plantedModule := new(PlantedModule)

	
	c.Bind(plantedModule)

	//"Move" all plant positions on up
	database, _ := sql.Open("sqlite3", "./rowa.db")
	database.Exec("UPDATE Plant SET PlantPosition = PlantPosition + 1 WHERE Harvested = 0 AND Module = ?", plantedModule.Module)

	//Add new  plant at pos 1
	statement, _ := database.Prepare("INSERT INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")
	statement.Exec(plantedModule.Module, 1, time.Now().Format("2006-01-02"), 0)

	rows, _ := database.Query("SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?", plantedModule.Module)

	var id int
	var ids []int
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids, id)
	}
	database.Exec("UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?", ids[0], plantedModule.Module)

	// Light off module
	if settings.ArduinoOn{
		go sensors.DeactivateModuleLight()
	}

	return
}
