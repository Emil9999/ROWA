package db

import (
	"log"
	"time"
)

type SensorData struct {
	Datetime       string  `json:"datetime"`
	Temp           float64 `json:"temperature"`
	LightIntensity float64 `json:"light_intensity"`
	Humidity       float64 `json:"humidity"`
	WaterLevel     float64 `json:"water_level"`   
	WaterTemp      float64 `json:"water_temp"`
    WaterpH        float64 `json:"water_ph"`
}

type PlantsPerPlantType struct {
	Name            string `json:"plant_type"`
	AvailablePlants int    `json:"available_plants"`
}

type BlinkModule struct {
	Module int    `json:"module"`
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
	case "plantable":
		for i := 1; i < 7; i++ {
			sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
			rows, err := store.Db.Query(sqlQuery, i)
			rows.Next()
			var id int
			rows.Scan(&id)
			
			if err != nil {
				log.Fatal(err)
			}
			rows.Close()
			
			sqlQuery = `SELECT PlantType, COUNT(PlantType) as AvailablePlants
						FROM Plant
								INNER JOIN Module M on Plant.Module = M.Id
								INNER JOIN PlantType PT on M.PlantType = PT.Name
						WHERE Harvested = 0
						AND M.Id = ?
						AND date(PlantDate, '+' || 7 || ' days') <= date('now')`
			
			rows, err = store.Db.Query(sqlQuery, i)
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
			
				if plantsPerPlantType.AvailablePlants-id == 0 && 6-id > 0 {
					p, found := find(plantsToHarvest, plantsPerPlantType.Name)
					if found {
						plantsToHarvest[p].AvailablePlants++
						log.Print(plantsPerPlantType)
					} else {
						plantsPerPlantType.AvailablePlants = 1
						plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)
						log.Print(plantsPerPlantType)
					}

				}

			}

		}
	case "modules":
		for i := 1; i < 7; i++ {
			sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
			rows, err := store.Db.Query(sqlQuery, i)
			rows.Next()
			var id int
			rows.Scan(&id)
			
			if err != nil {
				log.Fatal(err)
			}
			rows.Close()
			
			sqlQuery = `SELECT PlantType, COUNT(PlantType) as AvailablePlants
						FROM Plant
								INNER JOIN Module M on Plant.Module = M.Id
								INNER JOIN PlantType PT on M.PlantType = PT.Name
						WHERE Harvested = 0
						AND M.Id = ?
						AND date(PlantDate, '+' || 7 || ' days') <= date('now')`
			
			rows, err = store.Db.Query(sqlQuery, i)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				plantsPerPlantType := &PlantsPerPlantType{}
				err = rows.Scan(&plantsPerPlantType.Name, &plantsPerPlantType.AvailablePlants)
				if plantsPerPlantType.AvailablePlants-id == 0 && 6-id > 0 {
				plantsPerPlantType.AvailablePlants = i
				if err != nil {
					log.Fatal(err)
				}
				plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)}
				break

			}

		}
		/*sqlQuery = `SELECT PlantType, SUM(AvailableSpots) as AvailablePlants
		FROM Module
		where AvailableSpots > 0
		GROUP BY PlantType`*/
	default:
		log.Fatal("Wrong parameter passed into function")
	}
	return
}

func (store *Database) GetLastSensorEntry() (sensorData *SensorData, err error) {
	sqlQuery := `SELECT Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH
				 FROM SensorMeasurements
				 WHERE ID = (SELECT MAX(ID)  FROM SensorMeasurements)`

	row, err := store.Db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	sensorData = &SensorData{}

	row.Next()
	err = row.Scan(&sensorData.Datetime, &sensorData.Temp, &sensorData.LightIntensity, &sensorData.Humidity, &sensorData.WaterLevel, &sensorData.WaterTemp, &sensorData.WaterpH)

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
				WHERE M.Position = ? AND Harvested = 0
				ORDER BY PlantPosition`
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
		plantInfo = append(plantInfo, plantInfoPerModule)
	}
	
	return plantInfo, err

}

func find(slice []*PlantsPerPlantType, val string) (int, bool) {
	for i, item := range slice {
		if item.Name == val {
			return i, true
		}
	}
	return -1, false
}