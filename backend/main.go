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
)

func main() {
	if settings.Debug {
		setup.DbSetup()
	}
	sensors.Blink()
	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	//e.File("/", "index.html")
	e.GET("/dashboard/main", dashboard.GetDashInfo)

	e.POST("/plant/plant", plant.Plant)
	e.GET("/plant/available", utils.GetAvailableTypes)
	e.POST("/plant/finishedPlanting", plant.FinishPlanting)

	e.POST("/harvest/harvest", harvest.Harvest)
	e.GET("/harvest/available", utils.GetAvailableTypes)
	e.POST("/harvest/harvestdone", harvest.HarvestDone)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
