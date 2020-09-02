package sensor

import (
	"log"
	"time"
)

var ch = make(chan string, 100)

//var queue = list.New()

func WriteToCh(input string) {
	ch <- input
	//queue.PushBack(input)
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
		select {
		case v := <-ch:
			log.Println("Writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
		default:
			ReadSensorData(s)
		}
		/*for v := range ch {
			log.Println("Writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)

		}*/

		//ReadSensorData(s)
		//time.Sleep(time.Second)
	}
}
