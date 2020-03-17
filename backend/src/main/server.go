package main

import (
	"database/sql"
	"log"

	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"

	//"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/MarcelCode/ROWA/src/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database, err := sql.Open("sqlite3", "rowa.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStore(&db.Database{Db: database})

	if settings.Debug {
		db.FunctionStore.DbSetup()
	}

	if settings.ArduinoOn {
		go sensor.ReadSensorData()
		util.LightTimesRenew()
		go util.Runner()
	}

	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)

	e.GET("/harvest/get-plant", api.GetHarvestablePlantHandler)
	e.POST("/harvest/harvestdone", api.HarvestDoneHandler)

	e.GET("/plant/blinkstop", api.StopModuleBlink)
	e.GET("/plant/get-position", api.PlantHandler)
	e.POST("/plant/finish", api.FinishPlantingHandler)
	e.GET("/dashboard/cattree/:module", api.GetCatTreeDataHandler)

	e.POST("/adminSettings/insert-light", api.InsertLightTimes)
	e.GET("/adminSettings/get-light", api.GetLightTimes)
	e.POST("/adminSettings/changelight", api.ChangeLightState)

	e.GET("/adminSettings/get-types", api.GetPlantTypes)
	e.POST("/adminSettings/insertmodule-change", api.InsertModuleChanges)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

}
