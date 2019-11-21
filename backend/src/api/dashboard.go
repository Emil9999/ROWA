package api

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

func (store *dbStore) GetHarvestablePlants() (plantsToHarvest []*PlantsPerPlantType, err error) {
	sqlQuery := `SELECT PlantType, COUNT(PlantType) as AvailablePlantsPerPlantType
				FROM Plant
						 INNER JOIN Module M on Plant.Module = M.Id
						 INNER JOIN PlantType PT on M.PlantType = PT.Name
				where Harvested = 0
				  and date(PlantDate, '+' || GrowthTime || ' days') <= date('now')
				GROUP BY PlantType`

	rows, err := store.db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	//Iterating over the result and putting it into an array
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

// TODO Function for CatTree Information @Emil, @Behnaz

func (store *dbStore) GetLastSensorEntry() (sensorData *SensorData, err error) {
	sqlQuery := `SELECT Datetime, Temp, LightIntensity
				 FROM SensorMeasurements
				 ORDER BY Id DESC
				 LIMIT 1`

	row, err := store.db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	err = row.Scan(&sensorData.Datetime, &sensorData.Temp, &sensorData.LightIntensity)
	if err != nil {
		log.Fatal(err)
	}

	return sensorData, err
}
