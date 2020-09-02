package sensor

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

var ch = make(chan string, 100)
var queue = list.New()

func WriteToCh(input string) {
	//ch <- input
	queue.PushBack(input)
	log.Print("Wrote ", input, " to queue")
}

func ArduinoLoop() {
	s, err := SetupSerialConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	log.Println("Starting loop")
	_, err = s.Write([]byte("80"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("81"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("80"))
	time.Sleep(time.Second * 1)

	for {
		/*for v := range ch {
			log.Println("Writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)

		}*/
		for queue.Len() > 0 {
			v := queue.Front()
			str := fmt.Sprintf("%v", v)
			log.Println("Writing ", str, "to arduino")
			_, err = s.Write([]byte(str))
			if err != nil {
				log.Fatal(err)
			}
			queue.Remove(v)
			time.Sleep(2 * time.Second)

		}
		log.Println("Reading sensor data")
		ReadSensorData(s)
		time.Sleep(time.Second * 1)
	}
}
