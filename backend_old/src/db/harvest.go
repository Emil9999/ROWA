package db

import (
	"github.com/MarcelCode/ROWA/src/sensor"
	"fmt"
)

type PositionOnFarm struct {
	PlantPosition  int `json:"plant_position"`
	ModulePosition int `json:"module_position"`
}

type PositionOnFarm2 struct {
	PlantType      string `json:"plant_type" query:"plant_type"`
	PlantPosition  int    `json:"plant_position"`
	ModulePosition int    `json:"module_position"`
}

type Status struct {
	Message string `json:"status_message" query:"status_message"`
}

func (store *Database) HarvestDone(plantPosition *PositionOnFarm) (status *Status, err error) {
	status = &Status{}

	sqlQuery := `UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = ? AND Module= ?`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()

	_, err = statement.Exec(plantPosition.PlantPosition, plantPosition.ModulePosition)
	if err != nil {
		fmt.Println(err)
		status.Message = "error" + err.Error()
		return
	}
	
	sqlQuery = `UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position= ?`
	statement, _ = store.Db.Prepare(sqlQuery)
	_, err = statement.Exec(plantPosition.ModulePosition)
	if err != nil {
		fmt.Println(err)
		status.Message = "error" + err.Error()
		return
	}

	go sensor.BreathOffModule(plantPosition.ModulePosition)
	

	status.Message = "harvest done"
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
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	positionOnFarm = &PositionOnFarm{}
	rows.Next()
	err = rows.Scan(&positionOnFarm.PlantPosition, &positionOnFarm.ModulePosition)

	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (store *Database) GetAllHarvestablePlant() (positionsOnFarm []*PositionOnFarm2, err error) {

	sqlQuery := `SELECT PlantPosition, Module, PlantType
				 FROM Plant
						 INNER JOIN Module M on Plant.Module = M.Id
						 INNER JOIN PlantType PT on M.PlantType = PT.Name
				 where Harvested = 0
				  and date(PlantDate, '+' || GrowthTime || ' days') <= date('now')`

	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		positionOnFarm := &PositionOnFarm2{}

		err = rows.Scan(&positionOnFarm.PlantPosition, &positionOnFarm.ModulePosition, &positionOnFarm.PlantType)

		if err != nil {
			fmt.Println(err)
			return
		}
		positionsOnFarm = append(positionsOnFarm, positionOnFarm)

	}

	return
}

func (store *Database) MassHarvest(plantPositions []PositionOnFarm) (status *Status, err error) {
	status = &Status{}
	for _, plantPosition := range plantPositions {
		sqlQuery := `UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = ? AND Module= ?`
		statement, _ := store.Db.Prepare(sqlQuery)
		defer statement.Close()

		_, err = statement.Exec(plantPosition.PlantPosition, plantPosition.ModulePosition)
		if err != nil {
			fmt.Println(err)
			status.Message = "error"
			return
		}

		sqlQuery = `UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position= ?`
		statement, _ = store.Db.Prepare(sqlQuery)
		defer statement.Close()
		_, err = statement.Exec(plantPosition.ModulePosition)
		if err != nil {
			fmt.Println(err)
			status.Message = "error"
			return
		}

	}
	status.Message = "harvest done"
	return
}