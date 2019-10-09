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

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func plant(c echo.Context) (err error) {
	fmt.Println("This is server: request received")
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
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

	//Creating
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
}
