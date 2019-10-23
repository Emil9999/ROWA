package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type PlantType struct {
	Name string `json:"PlantType"`
}
type PositionOnFarm struct {
	PlantPosition  int `json: "plantPosition"`
	ModulePosition int `json: "modulePosition"`
}

func getDashInfo(c echo.Context) (err error) {
	database, _ := sql.Open("sqlite3", "./rowa.db")
	rows, _ := database.Query("SELECT date(PlantDate, '+'||GrowthTime||' days') as FinishPlantDate FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and FinishPlantDate <= date('now')")
	var availablePlant string
	var availablePlants []string
	for rows.Next() {
		rows.Scan(&availablePlant)
		availablePlants = append(availablePlants, availablePlant)
	}
	fmt.Println(availablePlants)
	return c.JSON(http.StatusOK, len(availablePlants))
}

func getAvailableTypes(c echo.Context) (err error) {
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
		plantTypes, _ := json.Marshal(plantTypes)
		fmt.Println(string(plantTypes))
		return c.JSON(http.StatusOK, string(plantTypes))
	} else {
		//TODO Make sure something useful is returned
		return
	}

}
func plant(c echo.Context) (err error) {
	fmt.Println("This is server: request received")

	plantType := new(PlantType)
	//Binding the received data to plantType
	c.Bind(plantType)

	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Selecting the modules that have available spots and match the plant type
	sqlQuery := fmt.Sprintf("SELECT Position FROM Module WHERE AvailableSpots>0 AND PlantType='%s'", plantType.Name)
	rows, _ := database.Query(sqlQuery)
	//We only need the position (?)
	var position int
	for rows.Next() {
		rows.Scan(&position)
		fmt.Println(strconv.Itoa(position))
	}
	//Returning only the last position since user can only plant in one module
	ret := c.JSON(http.StatusOK, position)
	return ret
}
func harvestDone(c echo.Context) (err error) {
	fmt.Println("Harvested!")
	plantPosition := new(PositionOnFarm)
	c.Bind(plantPosition)
	fmt.Println(plantPosition)
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")

	sqlQuery := fmt.Sprintf("UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = '%d' AND Module=%d", plantPosition.PlantPosition, plantPosition.ModulePosition)
	statement, _ := database.Prepare(sqlQuery)
	statement.Exec()

	return
}
func harvest(c echo.Context) (err error) {
	fmt.Println("Harvest request received")
	plantType := new(PlantType)
	//Binding the received data to plantType
	c.Bind(plantType)

	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")

	sqlQuery := fmt.Sprintf("SELECT PlantPosition, Module FROM Plant INNER JOIN Module M on Plant.Module = M.Id INNER JOIN PlantType PT on M.PlantType = PT.Name where Harvested = 0 and date(PlantDate, '+'||GrowthTime||' days') <= date('now') AND PlantType = '%s'", plantType.Name)

	rows, _ := database.Query(sqlQuery)
	//We only need the position (?)
	var position int
	var module int
	var arr []int
	for rows.Next() {
		rows.Scan(&position, &module)

		arr = append(arr, position)
		arr = append(arr, module)
		fmt.Println(strconv.Itoa(position) + " " + strconv.Itoa(module))
	}
	arr = arr[:2]
	ret := c.JSON(http.StatusOK, arr)
	return ret

}
func main() {
	dbSetup()
	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	//e.File("/", "index.html")
	e.GET("/dashboard/main", getDashInfo)
	e.POST("/plant/plant", plant)
	e.POST("/harvest/harvestdone", harvestDone)
	e.POST("/harvest/harvest", harvest)
	e.GET("/plant/available", getAvailableTypes)
	e.GET("/harvest/available", getAvailableTypes)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
func dbSetup() {
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Creating PlantType DB
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)") // TODO Why no ID?
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", 42) //TODO Why only one PlantType, whether we use two plants?
	statement.Exec("Basil", 21)
	rows, _ := database.Query("SELECT Name, Growthtime from PlantType")
	var name string
	var growthTime int
	for rows.Next() {
		rows.Scan(&name, &growthTime)
		fmt.Println(name + ": " + strconv.Itoa(growthTime))
	}
	//Creating Module DB
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Module (Id INTEGER PRIMARY KEY, PlantType TEXT, Position INTEGER UNIQUE, AvailableSpots INTEGER, TotalSpots INTEGER, FOREIGN KEY(PlantType) REFERENCES PlantType(Name))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT OR IGNORE INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec(1, "Basil", 6, 6)
	statement.Exec(2, "Basil", 4, 6)
	statement.Exec(3, "Lettuce", 0, 6)
	statement.Exec(4, "Lettuce", 1, 6)
	statement.Exec(5, "Lettuce", 0, 6)
	statement.Exec(6, "Basil", 0, 6)

	// Creating Plant DB
	statement, _ = database.Prepare("DROP TABLE IF EXISTS Plant")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Plant (Id INTEGER PRIMARY KEY, Module INTEGER, PlantPosition INTEGER, PlantDate TEXT, Harvested INTEGER, FOREIGN KEY (Module) REFERENCES Module(Id))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")

	// Add multiple Plants for testing purposes
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			statement.Exec(i+1, j+1, time.Now().Format("2006-01-02"), 0)
		}
	}

	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-09'  WHERE  Id = 1")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-08'  WHERE Id = 21")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-10'  WHERE Id = 22")
	statement.Exec()

	rows, _ = database.Query("SELECT Id, Position, PlantType, AvailableSpots, TotalSpots from Module")
	var Position int
	var Id int
	var PlantType string
	var AvailableSpots int
	var TotalSpots int
	for rows.Next() {
		rows.Scan(&Id, &Position, &PlantType, &AvailableSpots, &TotalSpots)
		fmt.Println(strconv.Itoa(Id), strconv.Itoa(Position)+PlantType+strconv.Itoa(AvailableSpots)+strconv.Itoa(TotalSpots))
	}
}
