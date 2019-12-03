package main

import (
	"api"
	"database/sql"
	"db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sensor"
	"settings"
)

func initialiseDb(filepath string) {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStore(&db.Database{Db: database})

	if settings.Debug {
		db.FunctionStore.DbSetup()
	}

	if settings.ArduinoOn {
		sensor.ReadSensorData()
	}
}

func initialiseEchoServer(enableCors bool) (e *echo.Echo) {
	e = echo.New()

	if enableCors {
		e.Use(middleware.CORS())
	}

	// Routes
	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)

	e.GET("/harvest/get-plant", api.GetHarvestablePlantHandler)
	e.POST("/harvest/harvestdone", api.HarvestDoneHandler)

	e.GET("/plant/get-position", api.PlantHandler)
	e.POST("/plant/finish", api.FinishPlantingHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

	return
}

func main() {
	initialiseDb("../../rowa.db")
	initialiseEchoServer(true)
}