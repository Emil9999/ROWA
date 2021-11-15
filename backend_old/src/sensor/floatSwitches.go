package sensor

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

func SetupFloater(pin string) gpio.PinIn {
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO" + pin)
	if p == nil {
		log.Fatal("Failed to find GPIO", pin)
	}

	if err := p.In(gpio.PullUp, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %s\n", p, p.Read())
	return p

}

func ReadFloater(pin string) {
	// Wait for rising edges (Low -> High) and print when one occur.
	p := SetupFloater(pin)
	for p.WaitForEdge(-1) {
		fmt.Printf("%s went %s\n", p, p.Read())
	}
	fmt.Println("terminated function")
}

func ReadFloaters(p_low, p_middle, p_top gpio.PinIn) float64 {

	if state(p_low) && state(p_middle) && !(state(p_top)) {
		return 100.0
	} else if state(p_low) && state(p_middle) {
		return 70.0
	} else if state(p_low) {
		return 30.0
	} else if !(state(p_low) && state(p_middle)) && state(p_top) {
		return 0.0
	} else {
		fmt.Println("Check floaters")
		return -1.0
	}

}

func state(p gpio.PinIn) bool {
	fmt.Println(p, p.Read())
	if p.Read() == gpio.High {
		return true
	} else {
		return false
	}
}
