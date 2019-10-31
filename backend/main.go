package main

import (
	"./dashboard"
	"./harvest"
	"./plant"
	"./settings"
	"./setup"
	"./utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"rowa/backend/sensors"
)

func main() {
	if settings.Debug {
		setup.DbSetup()
	}

	if settings.Raspberry {
		//Calling sensors Blink method as a goroutine
		go sensors.Blink()
	}

	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	e.GET("/dashboard/main", dashboard.GetDashInfo)
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
