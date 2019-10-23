package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"rowa/backend/settings"
	"strconv"
	"time"
)

type PlantType struct {
	Name string `json:"PlantType"`
}

type PlantedModule struct {
	Module int `json:"PlantedModule"`
}

func finishPlanting(c echo.Context) (err  error) {
	
	plantedModule := new(PlantedModule)

	c.Bind(plantedModule)

	//"Move" all plant positions on up
	database, _ := sql.Open("sqlite3", "./rowa.db")
	database.Exec("UPDATE Plant SET PlantPosition = PlantPosition + 1 WHERE Harvested = 0 AND Module = ?",plantedModule.Module)
	
	//Add new  plant at pos 1
	statement, _ := database.Prepare("INSERT INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")
	statement.Exec(plantedModule.Module, 1, time.Now().Format("2006-01-02"), 0)

	rows, _ := database.Query("SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?", plantedModule.Module)

	var id  int
	var ids []int
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids, id)
	}
	database.Exec("UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?",ids[0], plantedModule.Module)

 return
}

func getDashInfo(c echo.Context) (err error) {
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
	for rows.Next() {
		var row HarvestablePlantsPerType
		err = rows.Scan(&row.Name, &row.AvailablePlants)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, row)
	}
	fmt.Println(results)

	return c.JSON(http.StatusOK, results)
}

func getAvailableTypes(c echo.Context) (err error) {
	fmt.Println("Get request received")
	var plantTypes []string
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Getting all available plant types
	rows, _ := database.Query("SELECT DISTINCT PlantType FROM Module WHERE AvailableSpots>0")
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
func main() {
	if settings.Debug {
		dbSetup()
	}
	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	//e.File("/", "index.html")
	e.GET("/dashboard/main", getDashInfo)
	e.POST("/plant/plant", plant)
	e.POST("/plant/finishedPlanting", finishPlanting)
	e.GET("/plant/available", getAvailableTypes)
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
	statement, _ = database.Prepare("DROP TABLE IF EXISTS Module")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Module (Id INTEGER PRIMARY KEY, PlantType TEXT, Position INTEGER UNIQUE, AvailableSpots INTEGER, TotalSpots INTEGER, FOREIGN KEY(PlantType) REFERENCES PlantType(Name))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT OR IGNORE INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec(1, "Basil", 6, 6)
	statement.Exec(2, "Lettuce", 4, 6)
	statement.Exec(3, "Lettuce", 0, 6)
	statement.Exec(4, "Lettuce", 1, 6)
	statement.Exec(5, "Lettuce", 0, 6)
	statement.Exec(6, "Lettuce", 0, 6)

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

	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-09'  WHERE Id = 20")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-08'  WHERE Id = 21")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-10'  WHERE Id = 22")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET PlantDate = '2019-09-30'  WHERE Id = 4")
	statement.Exec()


	//Make  Empty Plant Pot
	statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 24")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 11")
	statement.Exec()
	statement, _ = database.Prepare("UPDATE Plant SET Harvested = 1  WHERE Id = 12")
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
