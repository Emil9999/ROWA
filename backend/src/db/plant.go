package db

import (
	"database/sql"
	"github.com/labstack/echo"
	"sensor"
	"settings"
	"time"
)

type PlantedModule struct {
	Module int `json:"planted_module" query:"planted_module"`
}

func (store *Database) Plant(plantType *PlantType) (position int, err error) {
	sqlQuery := `SELECT Position FROM Module WHERE AvailableSpots > 0 AND PlantType= ?`
	rows, err := store.Db.Query(sqlQuery)
	rows.Next()
	rows.Scan(&position)

	return
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
	if settings.ArduinoOn {
		go sensor.DeactivateModuleLight()
	}

	return
}

func (store *Database) FinishPlanting(plantedModule *PlantedModule) (status *Status, err error) {
	sqlQuery := `UPDATE Plant SET PlantPosition = PlantPosition + 1 WHERE Harvested = 0 AND Module = ?`
	_, err = store.Db.Exec(sqlQuery, plantedModule.Module)

	sqlQuery = `INSERT INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)`
	statement, _ := store.Db.Prepare(sqlQuery)
	_, err = statement.Exec(plantedModule.Module, 1, time.Now().Format("2006-01-02"), 0)

	sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
	rows, err := store.Db.Query(sqlQuery, plantedModule.Module)
	rows.Next()
	var id int
	rows.Scan(&id)

	sqlQuery = `UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?`
	_, err = store.Db.Exec(sqlQuery, id, plantedModule.Module)

	if settings.ArduinoOn {
		go sensor.DeactivateModuleLight()
	}

	status = &Status{Message: "Planting Done"}
	return
}
