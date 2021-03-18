package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/labstack/echo"
)

func LightOn(c echo.Context) (err error) {
	go sensor.Modules.Module1.lightOn()
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	go sensor.Modules.Module1.lightOff()
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	go sensor.Modules.Module1.breathOn()
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	go sensor.Modules.Module1.breathOff()
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	fmt.Println(sensor.Modules.module1.State)
	return c.JSON(http.StatusOK, "state")
}
