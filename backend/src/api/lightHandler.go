package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/labstack/echo"
)

func LightOn(c echo.Context) (err error) {
	go sensor.Modules.Module1.LightOn()
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	go sensor.Modules.Module1.LightOff()
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	go sensor.Modules.Module1.BreathOn()
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	go sensor.Modules.Module1.BreathOff()
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	fmt.Println(sensor.Modules.Module1.State)
	return c.JSON(http.StatusOK, "state")
}
