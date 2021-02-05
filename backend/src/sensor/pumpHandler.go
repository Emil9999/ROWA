package sensor

import "github.com/stianeikeland/go-rpio/v4"

func TriggerPumpX() {
	err := rpio.Open()
	if err != nil {
		Fatal.log(err)
	}
	pin := rpio.Pin(29)
	pin.Toggle()
}
