package main

import (
	"api"
	"database/sql"
	"db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// Intialise DB
	database, err := sql.Open("sqlite3", "./backend/rowa.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStore(&db.Database{Db: database})

	if api.Debug {
		db.FunctionStore.DbSetup()
	}

	// Echo instance
	e := echo.New()

	//Enabling CORS
	e.Use(middleware.CORS())

	// Routes
	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)
	//
	//e.POST("/plant/plant", api.Plant)
	//e.GET("/plant/available", api.GetAvailableTypes)
	//e.POST("/plant/finishedPlanting", api.FinishPlanting)

	//e.POST("/harvest/harvest", api.Harvest)
	//e.GET("/harvest/available", api.GetAvailableTypes)
	//e.POST("/harvest/harvestdone", api.HarvestDone)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
