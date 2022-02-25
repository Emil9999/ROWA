package sensor

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

func TriggerPump(state bool) {
	//trigger("17", state)

}

//Airstone obsolete in current version
func TriggerAirStone(state bool) {
	trigger("5", state)
}

func trigger(pin string, state bool) error {
	if _, err := driverreg.Init(); err != nil {
		log.Error(err)
	}
	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO" + pin)
	if p == nil {
		log.Error("Failed to find pin", pin)
	}
	fmt.Println(p.Read())
	if state {
		if err := p.Out(gpio.High); err != nil {
			log.Error(err)
			return err
		}
	} else {
		if err := p.Out(gpio.Low); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
