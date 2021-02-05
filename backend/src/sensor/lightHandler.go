package sensor

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func SwitchLight() {

}
func BlinkLight() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	pin := rpio.Pin(12)
	pin.Pwm()
	pin.Freq(5000)
	pin.DutyCycle(0, 32)
	// the LED will be blinking at 2000Hz
	// (source frequency divided by cycle length => 64000/32 = 2000)

	// five times smoothly fade in and out
	for i := 0; i < 10; i++ {
		for i := uint32(0); i < 32; i++ { // increasing brightness
			pin.DutyCycle(i, 32)
			time.Sleep(time.Second / 32)
		}
		for i := uint32(32); i > 0; i-- { // decreasing brightness
			pin.DutyCycle(i, 32)
			time.Sleep(time.Second / 32)
		}
	}
}
