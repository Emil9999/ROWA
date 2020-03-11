package sensor

import (
	"database/sql"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"strings"
	"time"
)

func setupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
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
