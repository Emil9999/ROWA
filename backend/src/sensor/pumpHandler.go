package sensor

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func TriggerPumpX() {
	err := rpio.Open()
	if err != nil {
		Fatal.log(err)
	}
	defer rpio.Close()
	pin.Output()
	pin := rpio.Pin(2)
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}