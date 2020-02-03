package db

import (
	"log"
)

type SensorData struct {
	Datetime       string  `json:"datetime"`
	Temp           float64 `json:"temperature"`
	LightIntensity float64 `json:"light_intensity"`
}

type PlantsPerPlantType struct {
	Name            string `json:"plant_type"`
	AvailablePlants int    `json:"available_plants"`
}

func (store *Database) GetPlantsPerType(farmAction string) (plantsToHarvest []*PlantsPerPlantType, err error) {
	sqlQuery := ``
	switch farmAction {
	case "harvestable":
		sqlQuery = `SELECT PlantType, COUNT(PlantType) as AvailablePlantsPerPlantType
					FROM Plant
							 INNER JOIN Module M on Plant.Module = M.Id
							 INNER JOIN PlantType PT on M.PlantType = PT.Name
					where Harvested = 0
					  and date(PlantDate, '+' || GrowthTime || ' days') <= date('now')
					GROUP BY PlantType`
	case "plantable":
		sqlQuery = `SELECT PlantType, SUM(AvailableSpots) as AvailablePlants
					FROM Module
					where AvailableSpots > 0
					GROUP BY PlantType`
	default:
		log.Fatal("Wrong parameter passed into function")
	}

	rows, err := store.Db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		plantsPerPlantType := &PlantsPerPlantType{}
		err = rows.Scan(&plantsPerPlantType.Name, &plantsPerPlantType.AvailablePlants)
		if err != nil {
			log.Fatal(err)
		}
		plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)
	}

	return
}

func (store *Database) GetLastSensorEntry() (sensorData *SensorData, err error) {
	sqlQuery := `SELECT Datetime, Temp, LightIntensity
				 FROM SensorMeasurements
				 WHERE ID = (SELECT MAX(ID)  FROM SensorMeasurements)`

	row, err := store.Db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	sensorData = &SensorData{}

	row.Next()
	err = row.Scan(&sensorData.Datetime, &sensorData.Temp, &sensorData.LightIntensity)

	if err != nil {
		log.Fatal(err)
	}

	return sensorData, err
}
