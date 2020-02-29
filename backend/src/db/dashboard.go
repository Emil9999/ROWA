package db

import (
	"fmt"
	"log"
	"time"
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

type PlantInfoPerModule struct {
	PlantType     string `json:"plant_type"`
	PlantPosition int    `json:"plant_position"`
	Age           int    `json:"age"`
	Harvestable   bool   `json:"harvestable"`
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

func (store *Database) GetCatTreeData(module int) (plantInfo []*PlantInfoPerModule, err error) {
	sqlQuery := `SELECT PlantType, PlantPosition, PlantDate, GrowthTime
				FROM Module M
				INNER JOIN Plant P ON M.Position = P.Module
				INNER JOIN PlantType PT ON M.PlantType = PT.Name
				WHERE M.Position = ? AND Harvested = 0`
	rows, err := store.Db.Query(sqlQuery, module)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		plantInfoPerModule := &PlantInfoPerModule{}
		var plantDate = ""
		var growthTime = 0

		err = rows.Scan(&plantInfoPerModule.PlantType, &plantInfoPerModule.PlantPosition, &plantDate, &growthTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(growthTime)
		layout := "2006-01-02"
		plantDateParsed, err := time.Parse(layout, plantDate)
		if err != nil {
			log.Fatal(err)
		}
		plantInfoPerModule.Age = int(time.Now().Sub(plantDateParsed).Hours() / 24)
		if plantInfoPerModule.Age >= growthTime {
			plantInfoPerModule.Harvestable = true
		} else {
			plantInfoPerModule.Harvestable = false
		}
		fmt.Println(plantInfoPerModule.Age)
		plantInfo = append(plantInfo, plantInfoPerModule)
	}

	return plantInfo, err

}
