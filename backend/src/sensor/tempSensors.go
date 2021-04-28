package sensor

import (
	"fmt"

	dht "github.com/MichaelS11/go-dht"
	"github.com/labstack/gommon/log"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

func InitDht(pin string) (*dht.DHT, error) {
	var dhtEmpty *dht.DHT
	setSensorPinsHigh(pin)
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)

		return dhtEmpty, err
	}

	sensor, err := dht.NewDHT(pin, dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		log.Print("DHT error")
		return dhtEmpty, err
	}

	return sensor, err
}
func ReadDht(outsideSensor *dht.DHT, boxSensor *dht.DHT) (map[string]float64, error) {

	values := make(map[string]float64)

	humidity, temperature, err := outsideSensor.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error outside:", err)
		return nil, err
	}
	values["humidity"] = humidity
	values["temperature"] = temperature

	boxHumidity, boxTemp, err := boxSensor.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error box:", err)
		return nil, err
	}

	values["humidity"] = boxHumidity
	values["temperature"] = boxTemp

	fmt.Println("humidity: %v\n", humidity)
	fmt.Println("temperature: %v\n", temperature)
	fmt.Println("boxhumidity: %v\n", boxHumidity)
	fmt.Println("temperature: %v\n", boxTemp)
	return values, nil

}

func setSensorPinsHigh(pin string) {
	if _, err := driverreg.Init(); err != nil {
		log.Error(err)
	}
	p1 := gpioreg.ByName(pin)
	if p1 == nil {
		log.Error("Failed to find pin")
	}
	if err := p1.Out(gpio.High); err != nil {
		log.Error(err)
	}
}
