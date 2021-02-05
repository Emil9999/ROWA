package sensor

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func BlinkLight(pin int64, toggle bool) {
	err := rpio.Open(pin)
	pin := rpio.Pin(pin)
	state := pin.Read()
	fmt.Println(state)
	//TODO put module light pins
	a := []int64{4, 22, 24}
	var b Blaster
	b.StartBlaster(a)

	b.ApplyBlaster(pin, 1)
	time.Sleep(time.Second)
	b.ApplyBlaster(pin, 0)

	for {
		if toggle {
			for i := 1; i < 100; i++ { // increasing brightness
				b.ApplyBlaster(pin, float64(i)/100)
				time.Sleep(time.Millisecond * 5)
			}
			for i := 100; i > 0; i-- { // decreasing brightness
				b.ApplyBlaster(pin, float64(i)/100)
				time.Sleep(time.Millisecond * 5)

			}
		}

	}
}
