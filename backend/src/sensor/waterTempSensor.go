package sensor

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadWaterTemp() {
	file, err := os.Open("/sys/bus/w1/devices/28-3c01d60727e3/w1_slave")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
