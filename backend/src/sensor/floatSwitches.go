package sensor

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

func setupFloater(pin string) gpio.PinIn {
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO", pin)
	if p == nil {
		log.Fatal("Failed to find GPIO", pin)
	}

	if err := p.In(gpio.PullDown, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %s\n", p, p.Read())
	return p

}

func ReadFloater(pin string) {
	// Wait for rising edges (Low -> High) and print when one occur.
	for setupFloater(pin).WaitForEdge(-1) {
		fmt.Printf("%s went %s\n", p, gpio.High)
	}
}
