package sensor

import (
	"database/sql"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"strings"
	"time"
	//"math"
)
type Times struct{
	TimeOn  int 
	TimeOff int 
}


func setupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "/dev/cu.usbmodem1414301", Baud: 9600}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ActivateModuleLight(moduleNumber int) {
	// Open Serial Connection
	s, err := setupSerialConnection()
	defer s.Close()

	// Create String from Module Number and send to connection
	moduleString := strconv.Itoa(moduleNumber)
	_, err = s.Write([]byte(moduleString))
	if err != nil {
		log.Fatal(err)
	}

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}

func DeactivateModuleLight() {
	// Open Serial Connection
	s, err := setupSerialConnection()
	defer s.Close()

	// Create String from Module Number and send to connection
	_, err = s.Write([]byte("99"))
	if err != nil {
		log.Fatal(err)
	}
	

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}

func LightSwitch(state bool) {

	s, err := setupSerialConnection()
	defer s.Close()

	//send turn off or on to arduino
	if (state){
		fmt.Println("d")
			_, err = s.Write([]byte("80"))
				
			//change DB light State
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 1")
			statement.Exec()
			database.Close()
	} else {
		fmt.Println("c")
		_, err = s.Write([]byte("81"))

			//change DB light State
			database, _ := sql.Open("sqlite3", "./rowa.db")
			statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 1")
			statement.Exec()
			database.Close()
	}
	
	if err != nil {
		log.Fatal(err)
	}
	

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}


func ReadSensorData() {
	var serialString string

	s, _ := setupSerialConnection()
	defer s.Close()

	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity) VALUES (?, ?, ?)")
	defer database.Close()

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		serialIncome := string(buf[:n])
		serialString += serialIncome

		if strings.HasSuffix(serialString, "\n") {
			datetime := time.Now().UTC().Format(time.RFC3339)
			raw_string := strings.TrimSuffix(serialString, "\r\n")
			data_array := strings.Split(raw_string, ",")
			temp, err1 := strconv.ParseFloat(data_array[0], 32)
			lightIntensity, err2 := strconv.ParseFloat(data_array[1], 32)

			if err1 == nil && err2 == nil {
				fmt.Println(datetime, temp, lightIntensity)
				statement.Exec(datetime, temp, lightIntensity)
			}

			serialString = ""
		}
	}
}
/*
func OnSwitchTimer(OnTime int){
	i, j, _ := time.Now().Clock()
	currentTime := i*60 +j
	timeOffset := (currentTime- OnTime)
	//TODO find other way to cast
	abstimeOffset := int(math.Abs(float64(timeOffset)))
	if timeOffset <= 0{
		time.Sleep(time.Duration(abstimeOffset) * time.Minute)
		
		} else {
		 time.Sleep(time.Duration(1440-abstimeOffset) * time.Minute)
		}
	for{
		LightSwitch(true)
		time.Sleep(24*time.Hour)
	}

}

func OffSwitchTimer(OffTime int){
	i, j, _ := time.Now().Clock()
	currentTime := i*60 +j
	timeOffset := (currentTime- OffTime)
	//TODO find other way to cast
	abstimeOffset := int(math.Abs(float64(timeOffset)))
	if timeOffset <= 0{
		time.Sleep(time.Duration(abstimeOffset) * time.Minute)
		
		} else {
		 time.Sleep(time.Duration(1440-abstimeOffset) * time.Minute)
		}
	for{
		LightSwitch(false)
		time.Sleep(24*time.Hour)
	}

}
*/
func LightTasker(){
	i, j, _ := time.Now().Clock()
			currentTime := i*60 +j
			fmt.Println(i, j, currentTime)
	
	
	ticker := time.NewTicker(1 * time.Minute)
	
	  for {
		select {
		  case <-ticker.C:
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

			i, j, _ := time.Now().Clock()
			currentTime := i*60 +j
			fmt.Println(restartTime.TimeOff, currentTime)
			offTimeOffset := (currentTime- restartTime.TimeOff)
			onTimeOffset := (currentTime- restartTime.TimeOn)
			//TODO find other way to cast
			rows.Close()
	
			if offTimeOffset == 0 {
				LightSwitch(false)
			
			} else if onTimeOffset == 0 {
				LightSwitch(true)
			}
						
		}
	}



}

