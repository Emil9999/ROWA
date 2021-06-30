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

	database, err := sql.Open("sqlite3", "./rowa.db")
	if err != nil {
		log.Error(err)
	}
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, ExternalTemp, BoxTemp, ExternalHumidity, WaterLevel, BoxHumidity) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()

	for {
		tempValues, err := ReadDht(outsideSensor)
		if err != nil {
			log.Error(err)
		}
		tempValuesBox, err := ReadDht(boxSensor)
		externalTemp := tempValues["temperature"]
		boxTemp := tempValuesBox["temperature"]
		externalHumidity := tempValues["humidity"]
		waterLevel := ReadWeight(hx711)
		boxHumidity := tempValuesBox["humidity"]
		datetime := time.Now().UTC().Format(time.RFC3339)
		fmt.Print(externalTemp, boxTemp, externalHumidity, boxHumidity, waterLevel, datetime)
		//Writing to local db
		statement.Exec(datetime, 1, 1, 1, 1, 1)
		time.Sleep(time.Second * 2)

	}

}
