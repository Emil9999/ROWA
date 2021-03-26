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
	"log"

	dht "github.com/d2r2/go-dht"
	"github.com/yryz/ds18b20"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/onewire"
	"periph.io/x/conn/v3/onewire/onewirereg"
	host "periph.io/x/host/v3"
)

func ReadTemp() (temp []byte) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}
	// Use onewirereg 1-wire bus registry to find the first available 1-wire bus.
	b, err := onewirereg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	//defer b.Close()
	// Prints out the gpio pin used.
	if p, ok := b.(onewire.Pins); ok {
		fmt.Printf("Q: %s", p.Q())
	}
	// Dev is a valid conn.Conn.
	d := &onewire.Dev{Addr: 23, Bus: b}

	// Send a command and expect a 5 bytes reply.
	write := []byte{}
	read := make([]byte, 5)
	if err := d.Tx(write, read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", read)
	return read
}
func ReadDht() {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 27, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}
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
		} else {
			fmt.Println(err)
		}
	}
	return
}
