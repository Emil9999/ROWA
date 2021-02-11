package sensor

import (
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
}
