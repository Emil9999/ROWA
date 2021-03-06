package api

import (
	"database/sql"
	"net/http"

	"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/echo"
)

func InsertLightTimes(c echo.Context) (err error) {
	//Insert new times
	//handle Data
	lightTimes := new(db.Times)

	err = c.Bind(lightTimes)

	status, err := db.FunctionStore.InsertLightTimes(lightTimes)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Light Time insertion failed")
	}
	return c.JSON(http.StatusOK, status)
}

func GetLightTimes(c echo.Context) (err error) {
	currentTimes, err := db.FunctionStore.GetLightTimes()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Couldnt get Light Times")
	}
	return c.JSON(http.StatusOK, currentTimes)
}

func ChangeLightState(c echo.Context) (err error) {

	state := new(db.LightState)

	err = c.Bind(state)
	if settings.ArduinoOn {
		if state.State == 0 {
			sensor.LightSwitch(false)
		} else {
			sensor.LightSwitch(true)
		}
	} else {
		if state.State == 0 {
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 1")
			statement.Exec()
			database.Close()
		} else {
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 1")
			statement.Exec()
			database.Close()
		}

	}
	if err != nil {
		return c.JSON(http.StatusNotFound, "LightSwitch Unsuccessfull")
	}
	return c.JSON(http.StatusOK, "OK")
}

func ChangePumpState(c echo.Context) (err error) {

	state := new(db.LightState)

	err = c.Bind(state)
	if settings.ArduinoOn {
		if state.State == 0 {
			sensor.TriggerPump(false)
		} else {
			sensor.TriggerPump(true)

		}
	} else {
		if state.State == 0 {
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 2")
			statement.Exec()
			database.Close()
		} else {
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 2")
			statement.Exec()
			database.Close()
		}

	}
	if err != nil {
		return c.JSON(http.StatusNotFound, "Pump Switch Unsuccessfull")
	}
	return c.JSON(http.StatusOK, "OK")
}

func ChangeAirState(c echo.Context) (err error) {

	state := new(db.LightState)

	err = c.Bind(state)
	if settings.ArduinoOn {
		if state.State == 0 {
			sensor.TriggerAirStone(false)
		} else {
			sensor.TriggerAirStone(true)
		}
	}
	if err != nil {
		return c.JSON(http.StatusNotFound, "Air Switch Unsuccessfull")
	}
	return c.JSON(http.StatusOK, "OK")
}

func GetPlantTypes(c echo.Context) (err error) {
	plantTypes, err := db.FunctionStore.GetPlantTypes()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Couldnt get Light Times")
	}
	return c.JSON(http.StatusOK, plantTypes)
}

func GetPumpTimes(c echo.Context) (err error) {
	pumpTime, err := db.FunctionStore.GetPumpTime()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Couldnt get Pump Times")
	}
	return c.JSON(http.StatusOK, pumpTime)
}

func InsertPumpTime(c echo.Context) (err error) {
	//Insert new times
	//handle Data
	pumpTime := new(db.PumpData)

	err = c.Bind(pumpTime)

	status, err := db.FunctionStore.InsertPumpTimes(pumpTime)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Pump time insertion failed")
	}
	return c.JSON(http.StatusOK, status)
}

func InsertModuleChanges(c echo.Context) (err error) {
	//Insert new times
	//handle Data
	plantTypes := new(db.PlantTypes)

	err = c.Bind(plantTypes)

	status, err := db.FunctionStore.InsertModuleChanges(plantTypes)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Module insertion failed")
	}
	return c.JSON(http.StatusOK, status)
}

func GetKnownPlantTypes(c echo.Context) (err error) {
	knowTypes, err := db.FunctionStore.GetKnownPlantTypes()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Couldnt get PlantTypes")
	}
	return c.JSON(http.StatusOK, knowTypes)
}

func RealityCheckHandler(c echo.Context) (err error){


	realitycheckData := new(db.RealityCheckData)

	err = c.Bind(realitycheckData)

	status, err := db.FunctionStore.RealityCheck(realitycheckData)
	if err != nil {
		return c.JSON(http.StatusNotFound, "RealityCheck insertion failed")
	}
	return c.JSON(http.StatusOK, status)

}