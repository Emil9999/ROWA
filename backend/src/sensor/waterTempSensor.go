package sensor

/*import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadWaterTemp() (temp string) {
	file, err := os.Open("/sys/bus/w1/devices/28-3c01d60727e3/w1_slave")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		val := temp[len(temp)-7 : len(temp)]
		fmt.Println(scanner.Text())
		fmt.Println(val)
		return temp
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}*/

import (
	"fmt"

	"github.com/yryz/ds18b20"
)

func ReadWaterTemp() (temp float64) {
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
		}
	}
	return
}
