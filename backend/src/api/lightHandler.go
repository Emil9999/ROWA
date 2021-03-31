package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/labstack/echo"
)

func LightOn(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.Module1.LightOn()
	case "2":
		go sensor.Modules.Module2.LightOn()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.Module1.LightOff()
	case "2":
		go sensor.Modules.Module2.LightOff()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.Module1.BreathOn()
	case "2":
		go sensor.Modules.Module2.BreathOn()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.Module1.BreathOff()
	case "2":
		go sensor.Modules.Module2.BreathOff()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		fmt.Println(sensor.Modules.Module1.State)
	case "2":
		fmt.Println(sensor.Modules.Module2.State)
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "state")
}
