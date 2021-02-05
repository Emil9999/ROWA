import "github.com/stianeikeland/go-rpio/v4"

func TriggerPumpX() {
	err := rpio.Open()
	pin := rpio.Pin(29)
	pin.Toggle()
}