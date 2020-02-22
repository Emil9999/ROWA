package util


import (

	"github.com/jasonlvhit/gocron"
	"sensor"
	"database/sql"
	"fmt"
	"time"
	"strconv"
	"strings"	

)
type Times struct{
	TimeOn  string 
	TimeOff string 
}

var light = gocron.NewScheduler()



func LightInit() {
	sqlQuery := `SELECT OnTime, OffTime
				 FROM TimeTable
				 WHERE ID = 1`
			database, _ := sql.Open("sqlite3", "./rowa.db")

			rows, err := database.Query(sqlQuery)

			

			if err != nil {
				return
			}

			restartTime := new(Times)
			restartTime = &Times{}
			rows.Next()
			err = rows.Scan(&restartTime.TimeOn, &restartTime.TimeOff)

			TimeOnArray := strings.Split(restartTime.TimeOn, ":")
			TimeOffArray := strings.Split(restartTime.TimeOff, ":")

			TimeOnHour, _ := strconv.Atoi(TimeOnArray[0])
			TimeOnMinute, _ := strconv.Atoi(TimeOnArray[1])

			TimeOffHour, _ := strconv.Atoi(TimeOffArray[0])
			TimeOffMinute, _ := strconv.Atoi(TimeOffArray[1])

			tOn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOnHour, TimeOnMinute, 0, 0, time.Local)
			tOff := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOffHour, TimeOffMinute, 0, 0, time.Local)
			
	light.Every(1).Day().At(restartTime.TimeOn).From(&tOn).Do(sensor.LightSwitch, true)
	light.Every(1).Day().At(restartTime.TimeOff).From(&tOff).Do(sensor.LightSwitch, false)
	
	
	rows.Close()
	_, time := light.NextRun()
	fmt.Println("OnTime " , time)





	
}
func Runner() {
	<-light.Start()
}

func LightRestart() (){
	light.Clear()
	light.Clear()

	sqlQuery := `SELECT OnTime, OffTime
				 FROM TimeTable
				 WHERE ID = 1`
			database, _ := sql.Open("sqlite3", "./rowa.db")

			rows, err := database.Query(sqlQuery)
			

			if err != nil {
				return
			}

			restartTime := new(Times)
			restartTime = &Times{}
			rows.Next()
			err = rows.Scan(&restartTime.TimeOn, &restartTime.TimeOff)
			
			TimeOnArray := strings.Split(restartTime.TimeOn, ":")
			TimeOffArray := strings.Split(restartTime.TimeOff, ":")

			TimeOnHour, _ := strconv.Atoi(TimeOnArray[0])
			TimeOnMinute, _ := strconv.Atoi(TimeOnArray[1])

			TimeOffHour, _ := strconv.Atoi(TimeOffArray[0])
			TimeOffMinute, _ := strconv.Atoi(TimeOffArray[1])

			tOn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOnHour, TimeOnMinute, 0, 0, time.Local)
			tOff := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOffHour, TimeOffMinute, 0, 0, time.Local)

			light.Every(1).Day().At(restartTime.TimeOn).From(&tOn).Do(sensor.LightSwitch, true)
			light.Every(1).Day().At(restartTime.TimeOff).From(&tOff).Do(sensor.LightSwitch, false)
	rows.Close()
}