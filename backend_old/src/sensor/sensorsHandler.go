package sensor

import (
	"database/sql"
	"fmt"
	"math/rand"
	"math"
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
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, ExternalTemp, BoxTemp, ExternalHumidity, WaterLevel, BoxHumidity) VALUES (?, ?, ?, ?, ?, ?)")
	defer database.Close()
	for {
		datetime := time.Now()
		externalTemp := roundTwoDecimals(rand.Float64()*1 + 22)
		boxHumidity := roundTwoDecimals(rand.Float64()*1 + 35)
		externalHumidity := roundTwoDecimals(rand.Float64()*10 + 30)
		waterLevel := 70
		boxTemp := roundTwoDecimals(rand.Float64()*1 + 22)
		datetimeStr := datetime.UTC().Format(time.RFC3339)
		fmt.Println(datetime, externalTemp, boxTemp, externalHumidity, waterLevel, boxHumidity)
		statement.Exec(datetimeStr, externalTemp, boxTemp, externalHumidity, waterLevel, boxHumidity)

		time.Sleep(2 * time.Second)
	}

}
func roundTwoDecimals(x float64) float64{
	return (math.Round(x*100)/100)
}
func ReadSensorData() {
	/*hx711, err := InitScale()
	if err != nil {
		log.Error(err)
	}*/
	p_low := SetupFloater("13")
	p_middle := SetupFloater("20")
	p_top := SetupFloater("21")
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
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, ExternalTemp, BoxTemp, ExternalHumidity, WaterLevel, BoxHumidity) VALUES (?, ?, ?, ?, ?, ?)")
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
		waterLevel := ReadFloaters(p_low, p_middle, p_top)
		boxHumidity := tempValuesBox["humidity"]
		datetime := time.Now().UTC().Format(time.RFC3339)
		fmt.Print(externalTemp, boxTemp, externalHumidity, boxHumidity, waterLevel, datetime)
		//Writing to local db
		statement.Exec(datetime, externalTemp, boxTemp, externalHumidity, waterLevel, boxHumidity)
		time.Sleep(time.Second * 2)

	}

}
