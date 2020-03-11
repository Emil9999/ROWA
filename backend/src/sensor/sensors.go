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


func setupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "COM5", Baud: 9600}
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
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
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
			humidity, err3 := strconv.ParseFloat(data_array[2], 32)
			waterLevel, err4 := strconv.ParseFloat(data_array[3], 32)
			waterTemp, err5 := strconv.ParseFloat(data_array[4], 32)
			waterpH, err6 := strconv.ParseFloat(data_array[5], 32)

			if err1 == nil && err2 == nil && err4 == nil && err3 == nil && err5 == nil && err6 == nil {
				fmt.Println(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
				statement.Exec(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
			}

			serialString = ""
		}
	}
}
