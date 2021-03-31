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
		go sensor.Modules.module1.lightOn()
	case "2":
		go sensor.Modules.module2.lightOn()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.module1.lightOff()
	case "2":
		go sensor.Modules.module2.lightOff()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.module1.breathOn()
	case "2":
		go sensor.Modules.module2.breathOn()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		go sensor.Modules.module1.breathOff()
	case "2":
		go sensor.Modules.module2.breathOff()
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	switch c.Param("id") {
	case "1":
		fmt.Println(sensor.Modules.module1.State)
	case "2":
		fmt.Println(sensor.Modules.module2.State)
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "state")
}
