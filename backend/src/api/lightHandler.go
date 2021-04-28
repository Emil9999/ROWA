package api

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/MarcelCode/ROWA/src/sensor"

	"github.com/labstack/echo"
)

func LightOn(c echo.Context) (err error) {
	var ModuleNum int
	ModuleNum, _ = strconv.Atoi(c.Param("id"))

	if (ModuleNum < 7 && ModuleNum > 0){
		sensor.LightOnModule(ModuleNum)
	} else {
		fmt.Println("Module not implemented")
	}
		
	
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	var ModuleNum int
	ModuleNum, _ = strconv.Atoi(c.Param("id"))

	if (ModuleNum < 7 && ModuleNum > 0){
		sensor.LightOffModule(ModuleNum)
	} else {
		fmt.Println("Module not implemented")
	}
		
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	var ModuleNum int
	ModuleNum, _ = strconv.Atoi(c.Param("id"))

	if (ModuleNum < 7 && ModuleNum > 0){
		sensor.BreathOnModule(ModuleNum)
	} else {
		fmt.Println("Module not implemented")
	}
		
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	var ModuleNum int
	ModuleNum, _ = strconv.Atoi(c.Param("id"))

	if (ModuleNum < 7 && ModuleNum > 0){
		sensor.BreathOffModule(ModuleNum)
	} else {
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
	case "3":
		fmt.Println(sensor.Modules.Module3.State)
	case "4":
		fmt.Println(sensor.Modules.Module4.State)
	case "5":
		fmt.Println(sensor.Modules.Module5.State)
	case "6":
		fmt.Println(sensor.Modules.Module6.State)
	default:
		fmt.Println("Module not implemented")
	}
	return c.JSON(http.StatusOK, "state")
}
