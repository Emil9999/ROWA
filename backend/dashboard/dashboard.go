package dashboard

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetDashInfo(c echo.Context) (err error) {
	// Query harvestable Plants per Plant type
	database, _ := sql.Open("sqlite3", "./rowa.db")
	rows, err := database.Query("SELECT PlantType, COUNT(PlantType) as AvailablePlantsPerPlantType FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and date(PlantDate, '+'||GrowthTime||' days') <= date('now') GROUP BY PlantType")

	if err != nil {
		log.Fatal(err)
	}

	type HarvestablePlantsPerType struct {
		Name            string `json:"plant_type"`
		AvailablePlants int    `json:"available_plants"`
	}

	var results []HarvestablePlantsPerType

	//Iterating over the result and putting it into an array
	for rows.Next() {
		var row HarvestablePlantsPerType
		err = rows.Scan(&row.Name, &row.AvailablePlants)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, row)
	}
	var breakRow HarvestablePlantsPerType
	breakRow.Name = "All Plants"
	breakRow.AvailablePlants = -1
	results = append(results, breakRow)
	rows, err = database.Query("SELECT PlantType, AvailableSpots FROM Module ")

	for rows.Next() {
		var row HarvestablePlantsPerType
		err = rows.Scan(&row.Name, &row.AvailablePlants)
		if err != nil {
			log.Fatal(err)
		}
		row.AvailablePlants = 6 - row.AvailablePlants
		results = append(results, row)
	}

	fmt.Println(results)
	return c.JSON(http.StatusOK, results)
}
