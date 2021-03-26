package sensor

import (
	"fmt"

	dht "github.com/MichaelS11/go-dht"
	"github.com/yryz/ds18b20"
)

func ReadDht() {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT("GPIO4", dht.Fahrenheit, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}

	humidity, temperature, err := dht.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("humidity: %v\n", humidity)
	fmt.Printf("temperature: %v\n", temperature)
}

func ReadTemp() (temp float64) {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
			return t
		} else {
			fmt.Println(err)
		}
	}
	return
}
