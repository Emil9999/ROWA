package api

import (
	"github.com/MarcelCode/ROWA/src/db"
	"database/sql"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/echo"
	"net/http"
	"github.com/MarcelCode/ROWA/src/sensor"
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
	currentTimes, err :=  db.FunctionStore.GetLightTimes()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Couldnt get Light Times")
	}
	return c.JSON(http.StatusOK, currentTimes)
}

func ChangeLightState(c echo.Context) (err error) {
	
	state := new(db.LightState)

	err = c.Bind(state)
	if settings.ArduinoOn{
	if state.State == 0{
	sensor.LightSwitch(false)
	} else {
		sensor.LightSwitch(true)
	}
} else { 
	if state.State == 0{
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
