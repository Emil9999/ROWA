package sensor

import (
	"fmt"
	"log"

	dht "github.com/d2r2/go-dht"
	"github.com/yryz/ds18b20"
)

func ReadDht() {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
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
