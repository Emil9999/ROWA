package utils

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type PlantType struct {
	Name string `json:"PlantType"`
}

func GetAvailableTypes(c echo.Context) (err error) {
	fmt.Println("Get request received")
	var plantTypes []string
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Getting all available plant types
	rows, _ := database.Query("")
	if c.Path() == "/plant/available" {
		rows, _ = database.Query("SELECT DISTINCT PlantType FROM Module WHERE AvailableSpots>0")

	} else {
		rows, _ = database.Query("SELECT PlantType FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and date(PlantDate, '+'||GrowthTime||' days') <= date('now') GROUP BY PlantType")
	}

	var plantType string
	//Iterating over the result and putting it them into an array
	for rows.Next() {
		rows.Scan(&plantType)
		plantTypes = append(plantTypes, plantType)
	}

	//Checking if nothing is returned
	if len(plantTypes) > 0 {
		fmt.Println(len(plantTypes))
		return c.JSON(http.StatusOK, plantTypes)
	} else {
		//TODO Make sure something useful is returned
		return
	}

}
