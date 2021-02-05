package main

import (
	"database/sql"
	"log"

	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	database, err := sql.Open("sqlite3", "rowa.db")
	if err != nil {
		log.Println(err)
	}
	err = settings.LoadConfiguration("settings.json")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStore(&db.Database{Db: database})

	if settings.Debug {
		db.FunctionStore.DbSetup()
	}

	//defer s.Close()
	//go sensor.ReadFakeSensorData()

	/*if err != nil {
		log.Print("No arduino found, faking data..")
		go sensor.ReadFakeSensorData()
	} else {
		log.Print("Arduino found..")
		go sensor.ArduinoLoop(s)
		go sensor.ReadSensorData(s)
		util.LightTimesRenew()
		util.PumpTimesRenew()
		go util.Runner()
	}*/
	//go sensor.TriggerPumpX()
	go sensor.BlinkLight(17)
	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)
	e.GET("/dashboard/plantable-modules", api.GetPlantableModulesHandler)
	e.POST("/dashboard/blink", api.StartBlink)

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
	e.GET("/adminSettings/get-knowntypes", api.GetKnownPlantTypes)
	e.POST("/adminSettings/insertmodule-change", api.InsertModuleChanges)

	e.POST("/adminSettings/insert-pump", api.InsertPumpTime)
	e.GET("/adminSettings/get-pump", api.GetPumpTimes)
	e.POST("/adminSettings/changePump", api.ChangePumpState)
	e.POST("/adminSettings/changeAir", api.ChangeAirState)
	e.POST("/admin/reality-check", api.RealityCheckHandler)

	e.GET("/plant/get-all", api.AllPlantHandler)
	e.POST("/plant/plant-all", api.MassPlantingHandler)

	e.GET("/harvest/get-all", api.AllHarvestHandler)
	e.POST("/harvest/harvest-all", api.MassHarvestHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

}
