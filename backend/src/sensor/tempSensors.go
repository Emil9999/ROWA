package sensor

import (
	"fmt"

	dht "github.com/MichaelS11/go-dht"
)

func ReadDht() (map[string]float64, error) {
	values := make(map[string]float64)
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return nil, err
	}

	outsideSensor, err := dht.NewDHT("GPIO22", dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return nil, err
	}
	boxSensor, err := dht.NewDHT("GPIO27", dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return nil, err
	}

	humidity, temperature, err := outsideSensor.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
		return nil, err
	}
	values["humidity"] = humidity
	values["temperature"] = temperature

	boxHumidity, boxTemp, err := boxSensor.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
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
