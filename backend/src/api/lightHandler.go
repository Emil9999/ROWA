package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/labstack/echo"
)

func LightOn(c echo.Context) (err error) {
	module := sensor.Modules.Modules[c.Param("id")]
	go module.LightOn()
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	module := sensor.Modules.Modules[c.Param("id")]
	go module.LightOff()
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	param := c.Param("id")
	module := sensor.Modules.Modules[param]
	go module.BreathOn()
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	param := c.Param("id")
	fmt.Println(param)
	module := sensor.Modules.Modules[param]
	go module.BreathOff()
	//go sensor.Modules.Module1.BreathOff("id")
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	//fmt.Println(sensor.Modules.Module1.State)
	return c.JSON(http.StatusOK, "state")
}
