package db

import (
	"sensor"
	"settings"
)

type PositionOnFarm struct {
	plantPosition  int `json: "plant_position" query: "plant_position"`
	modulePosition int `json: "module_position" query: "module_position"`
}

type Status struct {
	message string `json:"status_message" query: status_message`
}

func (store *Database) HarvestDone(plantPosition *PositionOnFarm) (status *Status, err error) {
	status = &Status{}

	sqlQuery := `UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = ? AND Module= ?`
	statement, _ := store.Db.Prepare(sqlQuery)
	_, err = statement.Exec(plantPosition.plantPosition, plantPosition.modulePosition)
	if err != nil {
		status.message = "error"
		return
	}

	sqlQuery = `UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position= ?`
	statement, _ = store.Db.Prepare(sqlQuery)
	_, err = statement.Exec(plantPosition.modulePosition)
	if err != nil {
		status.message = "error"
		return
	}

	if settings.ArduinoOn {
		go sensor.DeactivateModuleLight()
	}

	status.message = "harvest done"
	return
}

func (store *Database) GetHarvestablePlant(plantType *PlantType) (positionOnFarm *PositionOnFarm, err error) {

	sqlQuery := `SELECT PlantPosition, Module
				 FROM Plant
						 INNER JOIN Module M on Plant.Module = M.Id
						 INNER JOIN PlantType PT on M.PlantType = PT.Name
				 where Harvested = 0
				  and date(PlantDate, '+' || GrowthTime || ' days') <= date('now')
				  AND PlantType = ?`

	rows, err := store.Db.Query(sqlQuery, plantType.Name)
	if err != nil {
		return
	}

	positionOnFarm = &PositionOnFarm{}
	rows.Next()
	err = rows.Scan(&positionOnFarm.plantPosition, &positionOnFarm.modulePosition)

	if err != nil {
		return
	}

	if settings.ArduinoOn {
		go sensor.ActivateModuleLight(positionOnFarm.modulePosition)
	}

	return
}
