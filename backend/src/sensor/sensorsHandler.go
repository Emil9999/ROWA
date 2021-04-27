package sensor

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/labstack/gommon/log"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/spi/spireg"
)

/*Serial Port Configs
/dev/ttyACM0 - Rasp
/dev/cu.usbmodem1434301 Macbook
COM5 windows
*/

/*func SetupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Print(err)
		return
	}

	return s, err
}

func ActivateModuleLight(moduleNumber int) {
	// Create String from Module Number and send to connection
	moduleString := strconv.Itoa(moduleNumber)

	// Give Connection time to send Data
	WriteToChannel(moduleString)
}
func TriggerPump(state bool) {

	if state {
		WriteToChannel("90")
		fmt.Println("Pump on")
	} else {
		WriteToChannel("91")
		fmt.Println("Pump off")
	}

}

func TriggerAirStone(state bool) {
	if state {
		WriteToChannel("70")
		fmt.Println("Airstone on")
	} else {
		WriteToChannel("71")
		fmt.Println("Airstone off")
	}
}

func DeactivateModuleLight() {
	WriteToChannel("99")
}

func LightSwitch(state bool) {
	fmt.Println("Light switch triggered")

	//send turn off or on to arduino

	if state {
		WriteToChannel("80")
		fmt.Println("Light on")
		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 1")
		statement.Exec()
		database.Close()
	} else {
		WriteToChannel("81")
		fmt.Println("Light off")
		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 1")
		statement.Exec()
		database.Close()
	}

}*/
func DetectRpi() bool {
	if _, err := driverreg.Init(); err != nil {
		log.Error(err)
	}

	// Using SPI as an example. See package ./spi/spireg for more details.
	p, err := spireg.Open("")
	if err != nil {
		return false
	}
	defer p.Close()
	return true
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
func ReadSensorData() {
	hx711, err := InitScale()
	if err != nil {
		log.Error(err)
	}
	outsideSensor, err := InitDht("GPIO22")
	if err != nil {
		log.Error(err)
	}
	boxSensor, err := InitDht("GPIO27")
	if err != nil {
		log.Error(err)
	}

	/*database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()*/

	for {
		tempValues, err := ReadDht(outsideSensor, boxSensor)
		if err != nil {
			log.Error(err)
		}
		externalTemp := tempValues["temperature"]
		boxTemp := tempValues["boxTemp"]
		externalHumidity := tempValues["humidity"]
		waterLevel := ReadWeight(hx711)
		boxHumidity := tempValues["boxHumidity"]
		datetime := time.Now().UTC().Format(time.RFC3339)
		fmt.Print(externalTemp, boxTemp, externalHumidity, boxHumidity, waterLevel, datetime)
		//Writing to local db
		//statement.Exec(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
		time.Sleep(time.Second)

	}

}
