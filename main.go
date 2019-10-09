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
	Name       string `json:"name"`
	GrowthTime int    `json:"GrowthTime"`
}

func plant(c echo.Context) (err error) {
	fmt.Println("This is server: request received")
	u := new(PlantType)

	ret := c.JSON(http.StatusOK, u)
	fmt.Println(u)
	return ret
}
func main() {
	dbSetup()
	// Echo instance
	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	//e.File("/", "index.html")
	e.POST("/plant/greet", plant)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
func dbSetup() {
	//Opening DB
	database, _ := sql.Open("sqlite3", "./rowa.db")

	//Creating PlantType DB
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Id INTEGER PRIMARY KEY, Name TEXT, GrowthTime INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", "42")
	rows, _ := database.Query("SELECT Id, Name, Growthtime from PlantType")
	var name string
	var growthTime int
	var id int
	for rows.Next() {
		rows.Scan(&id, &name, &growthTime)
		fmt.Println(name + ": " + strconv.Itoa(growthTime))
	}
	//Creating Module DB
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS Module (Id INTEGER PRIMARY KEY, PlantType TEXT FORIEN KEY REFERENCES PlantType(Id), Position INTEGER, AvailableSpots INTEGER, TotalSpots INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec("1", "0", "6", "6")
	rows, _ = database.Query("SELECT Position, PlantType, AvailableSpots, TotalSpots from Module")
	var Position int
	var PlantType string
	var AvailableSpots int
	var TotalSpots int
	for rows.Next() {
		rows.Scan(&Position, &PlantType, &AvailableSpots, &TotalSpots)
		fmt.Println(strconv.Itoa(Position) + strconv.Itoa(AvailableSpots) + strconv.Itoa(TotalSpots))
	}
}
