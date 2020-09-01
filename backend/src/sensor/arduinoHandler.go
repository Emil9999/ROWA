package sensor

import (
	"fmt"
	"log"
	"time"
)

var ch = make(chan string, 100)

func WriteToCh(input string) {
	ch <- input
}

func ArduinoLoop() {
	s, err := SetupSerialConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	for {
		for v := range ch {
			fmt.Println("writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)

		}
		ReadSensorData(s)
		time.Sleep(time.Second * 2)
	}
}
