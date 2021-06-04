package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/util"
	log "github.com/labstack/gommon/log"

	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	host "periph.io/x/host/v3"
)

func main() {
	database, err := sql.Open("sqlite3", "rowa.db")
	if err != nil {
		fmt.Println(err)
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
	host.Init()

	//defer s.Close()

	if sensor.DetectRpi() {
		log.Print("Raspi found..")
		//go sensor.ReadSensorData()

		//sensor.SetSensorPinsHigh()
		go sensor.TestDht()
		go sensor.TriggerAirStone()
		time.Sleep(time.Second)
		go sensor.TriggerPump()

		sensor.SetupLight()
		//sensor.LightAllOff()
		time.Sleep(time.Second)
		//sensor.LightAllOn()
		hx711, err := InitScale()
		if err != nil {
			log.Error(err)
		}
		util.LightTimesRenew()
		util.PumpTimesRenew()
		go util.Runner()
	} else {
		log.Print("No raspi found, faking data..")
		//go sensor.ReadFakeSensorData()
	}
	//sensor.TriggerAirStone()
	waterLevel := ReadWeight(hx711)
	//go sensor.TriggerPumpX()
	//go sensor.BlinkLight(17, true)
	//go sensor.ReadWaterTemp()
	//go sensor.ReadTemp()
	//go sensor.ReadWeight()
	//go sensor.TestScale()
	//c := sensor.Spiinit()
	/*sensor.TriggerAirStone()
	time.Sleep(time.Second * 2)
	sensor.TriggerAirStone()
	fmt.Println("pump")
	sensor.TriggerPump()
	sensor.TriggerPump()*/
	//time.Sleep(time.Second * 2)

	//sensor.LightOnModule(5)
	//time.Sleep(time.Second * 5)
	///sensor.LightOffModule(5)

	//sensor.LightAllOn()
	//time.Sleep(time.Second * 5)
	//sensor.BreathOnModule(5)

	//sensor.BreathOnModule(2)
	///time.Sleep(time.Second * 30)
	//sensor.BreathOffModule(5)
	//sensor.BreathOnModule(6)
	//time.Sleep(time.Second * 30)

	//sensor.BreathOffModule(6)

	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	e.GET("/light/:id/on", api.LightOn)
	e.GET("/light/:id/off", api.LightOff)
	e.GET("/light/:id/breath-on", api.BreathOn)
	e.GET("/light/:id/breath-off", api.BreathOff)
	e.GET("/light/:id/state", api.State)

	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)
	e.GET("/dashboard/plantable-modules", api.GetPlantableModulesHandler)
	//e.POST("/dashboard/blink", api.StartBlink)

	e.GET("/harvest/get-plant", api.GetHarvestablePlantHandler)
	e.POST("/harvest/harvestdone", api.HarvestDoneHandler)

	//e.GET("/plant/blinkstop", api.StopModuleBlink)
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
