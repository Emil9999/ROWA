package sensor

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadWaterTemp() (temp string, err error) {
	file, err := os.Open("/sys/bus/w1/devices/28-3c01d60727e3/w1_slave")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		return temp, _
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return _, err
	}

}
