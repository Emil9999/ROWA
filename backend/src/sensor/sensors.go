package sensor

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
	//"github.com/tarm/serial"
)

/*Serial Port Configs
/dev/ttyACM0 - Rasp
/dev/cu.usbmodem1434301 Macbook
COM5 windows
*/

func SetupSerialConnection() (port io.ReadWriteCloser, err error) {
	fmt.Println("Setting serial connection")
	/*c := &serial.Config{Name: "COM5", Baud: 9600, ReadTimeout: time.Second * 1}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Print(err)
	}
	return*/

	// Set up options.
	options := serial.OpenOptions{
		PortName:        "/dev/ttyACM0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err = serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	// Write 4 bytes to the port.
	b := []byte("80")
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
	return
}

func ActivateModuleLight(moduleNumber int) {
	// Open Serial Connection
	s, err := SetupSerialConnection()
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
func TriggerPump(state bool) {
	// Open Serial Connection
	s, err := SetupSerialConnection()
	defer s.Close()
	/*
		// Create String from Module Number and send to connection
		moduleString := strconv.Itoa(moduleNumber)
		_, err = s.Write([]byte(moduleString))
		if err != nil {
			log.Fatal(err)
		}

		// Give Connection time to send Data
		time.Sleep(2 * time.Second)*/
	if state {
		_, err = s.Write([]byte("90"))
		fmt.Println("true")
	} else {
		_, err = s.Write([]byte("91"))
		fmt.Println("false")
	}

	if err != nil {
		log.Fatal(err)
	}

	// Give Connection time to send Data
	time.Sleep(1 * time.Second)
}

func DeactivateModuleLight() {
	// Open Serial Connection
	s, err := SetupSerialConnection()
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
	fmt.Println("Light switch triggered")
	s, err := SetupSerialConnection()

	//send turn off or on to arduino
	if state {
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
	defer s.Close()

}
func ReadFakeSensorData() {
	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()
	for {
		datetime := time.Now()
		temp := rand.Float32()*1 + 22
		lightIntensity := rand.Float32()*1 + 35
		humidity := rand.Float32()*10 + 30
		waterLevel := rand.Float32()*4 + 19
		waterTemp := rand.Float32()*1 + 22
		waterpH := rand.Float32() + 6
		datetimeStr := datetime.UTC().Format(time.RFC3339)
		fmt.Println(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
		statement.Exec(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)

		time.Sleep(2 * time.Second)
	}

}
func ReadSensorData() (err error) {
	var serialString string
	s, err := SetupSerialConnection()
	if err != nil {
		return err
	}

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
			datetime := time.Now()
			raw_string := strings.TrimSuffix(serialString, "\r\n")
			data_array := strings.Split(raw_string, ",")
			if len(data_array) >= 6 {
				temp, err1 := strconv.ParseFloat(data_array[0], 32)
				lightIntensity, err2 := strconv.ParseFloat(data_array[1], 32)
				humidity, err3 := strconv.ParseFloat(data_array[2], 32)
				waterLevel, err4 := strconv.ParseFloat(data_array[3], 32)
				waterTemp, err5 := strconv.ParseFloat(data_array[4], 32)
				waterpH, err6 := strconv.ParseFloat(data_array[5], 32)
				if err1 == nil && err2 == nil && err4 == nil && err3 == nil && err5 == nil && err6 == nil {
					fmt.Println(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
					datetime := datetime.UTC().Format(time.RFC3339)
					//Writing to local db
					statement.Exec(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
				}

			}
			serialString = ""
		}
		defer s.Close()

	}

}
