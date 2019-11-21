package main

import (
	"api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if api.Debug {
		api.DbSetup()
	}
	// Echo instance
	e := echo.New()
	//Enabling CORS
	e.Use(middleware.CORS())
	// Routes
	//e.File("/", "index.html")
	//e.GET("/dashboard/main", api.GetDashInfo)
	//
	//e.POST("/plant/plant", api.Plant)
	//e.GET("/plant/available", api.GetAvailableTypes)
	//e.POST("/plant/finishedPlanting", api.FinishPlanting)

	e.POST("/harvest/harvest", api.Harvest)
	//e.GET("/harvest/available", api.GetAvailableTypes)
	//e.POST("/harvest/harvestdone", api.HarvestDone)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
