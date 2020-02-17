package api

import (
	"db"
	
	"github.com/labstack/echo"
	"net/http"
	"sensor"
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
	if state.State == 0{
	sensor.LightSwitch(false)
	} else {
		sensor.LightSwitch(true)
	}
	if err != nil {
		return c.JSON(http.StatusNotFound, "LightSwitch Unsuccessfull")
	}
	return c.JSON(http.StatusOK, "OK")
}
