package api
/*
import (
	"db"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

*/

func InsertLightTimes(c echo.Context) (err error) {
//Insert new times
//handle Data
return c.JSON(http.StatusOK, status)
}




func GetLightTimes(c echo.Context) (err error) {
	//get new times
	//handle Data
	return c.JSON(http.StatusOK, currentTimes)
}

func ChangeLightState(c echo.Context) (err error) {
	//sensor.SwitchLight(bool)
	return c.JSON(http.StatusOK, status)
}
