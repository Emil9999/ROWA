package main

import (
	"./dashboard"
	"./harvest"
	"./plant"
	"./sensors"
	"./settings"
	"./setup"
	"./utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if settings.Debug {
		setup.DbSetup()
	}

	if settings.ArduinoOn{
		go sensors.ReadSensorData()
	}

	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	e.GET("/dashboard/harvestable", dashboard.GetHarvestablePlants)
	e.GET("/dashboard/all", dashboard.GetAllPlants)
	e.GET("/dashboard/sensor-data", dashboard.GetSensorInfo)
	e.GET("/plant/available", utils.GetAvailableTypes)
	e.GET("/harvest/available", utils.GetAvailableTypes)

	e.POST("/plant/plant", plant.Plant)
	e.POST("/plant/finishedPlanting", plant.FinishPlanting)
	e.POST("/harvest/harvest", harvest.Harvest)
	e.POST("/harvest/harvestdone", harvest.HarvestDone)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
