package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type PlantType struct {
	Name string `json:"PlantType"`
}

func getAvailableTypes(c echo.Context) (err error) {
	fmt.Println("Get request received")
	return
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
	dbSetup()
	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	//e.File("/", "index.html")
	e.POST("/plant/plant", plant)
	e.GET("/plant/available", getAvailableTypes)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
func dbSetup() {
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")
	//Creating PlantType DB
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", "42")
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
	statement.Exec(1, "Lettuce", 6, 6)
	statement.Exec(2, "Lettuce", 4, 6)
	statement.Exec(3, "Lettuce", 0, 6)
	statement.Exec(4, "Lettuce", 1, 6)
	statement.Exec(5, "Lettuce", 0, 6)
	statement.Exec(6, "Lettuce", 1, 6)

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
