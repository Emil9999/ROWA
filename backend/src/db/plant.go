package db

import (
	"errors"
	"github.com/labstack/gommon/log"
	"sensor"
	"settings"
	"time"
)

type PlantedModule struct {
	Module int `json:"planted_module" query:"planted_module"`
}

func (store *Database) Plant(plantType *PlantType) (modulePosition int, err error) {
	sqlQuery := `SELECT Position FROM Module WHERE AvailableSpots > 0 AND PlantType= ?`
	rows, err := store.Db.Query(sqlQuery, plantType.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&modulePosition)
		if err != nil {
			return
		}
	} else {
		err = errors.New("no data available")
		return
	}

	if settings.ArduinoOn {
		go sensor.ActivateModuleLight(modulePosition)
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
	defer rows.Close()
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
