package sensor

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

func TriggerPump() {
	trigger("17")

}
func TriggerAirStone() {
	trigger("4")
}
func trigger(pin string) error {
	if _, err := driverreg.Init(); err != nil {
		log.Error(err)
	}
	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO" + pin)
	if p == nil {
		log.Error("Failed to find pin", pin)
	}
	fmt.Println(p.Read())
	//if
	// Set the pin as output High.
	/*if err := p.Out(gpio.High); err != nil {
		log.Fatal(err)
	}*/
	return nil
}
