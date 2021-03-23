package sensor

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"periph.io/x/conn/v3"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	host "periph.io/x/host/v3"
)

type ModulesStruct struct {
	Modules map[string]Module
	c       conn.Conn
}

type Module struct {
	Pin           int
	State         bool
	StopBreathing chan bool
	BreathState   bool
}

var (
	globalIntensity  = 255
	breathingSpeedMs = 10
)

func (lm *Module) init() {
	p := gpioreg.ByName("GPIO" + strconv.Itoa(lm.Pin))
	if p == nil {
		log.Fatal("Failed to find GPIO" + strconv.Itoa(lm.Pin))
	}
	if err := p.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

}

func (ms *ModulesStruct) SetPinsHigh(pin int) {
	fmt.Println("called pin", pin)
	for _, module := range ms.Modules {
		if module.BreathState {
			module.StopBreathing <- true
		}
		if module.Pin != pin {

			fmt.Println("high", module.Pin)
			p := gpioreg.ByName("GPIO" + strconv.Itoa(module.Pin))
			if err := p.Out(gpio.High); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("low", module.Pin)

			p := gpioreg.ByName("GPIO" + strconv.Itoa(module.Pin))
			if err := p.Out(gpio.Low); err != nil {
				log.Fatal(err)
			}
		}

	}
}

func (lm *Module) LightOn() {
	Modules.SetPinsHigh(lm.Pin)
	writeToPoti(globalIntensity)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) LightOff() {
	Modules.SetPinsHigh(lm.Pin)
	writeToPoti(0)
	lm.State = false
	fmt.Println("State", lm.State)
}

func (lm *Module) BreathOn() {
	fmt.Println("Start breathing")
	fmt.Println("State", lm.State)
	intensityDown := lm.State
	var intensity int
	if lm.State {
		intensity = globalIntensity
	} else {
		intensity = 0
	}
	Modules.SetPinsHigh(lm.Pin)

	for {
		select {
		case <-lm.StopBreathing:
			fmt.Println("Stop infinite loop")
			return
		default:
			if intensityDown {
				intensity--
			} else {
				intensity++
			}

			if intensity == globalIntensity || intensity == 0 {
				intensityDown = !intensityDown
			}
			writeToPoti(intensity)
			time.Sleep(time.Duration(breathingSpeedMs) * time.Millisecond)
		}
	}
	lm.BreathState = true
}

func (lm *Module) BreathOff() {
	fmt.Println("Stop breathing")
	fmt.Println("State", lm.State)
	lm.StopBreathing <- true
	Modules.SetPinsHigh(lm.Pin)

	var intensity int
	if lm.State {
		intensity = globalIntensity
	} else {
		intensity = 0
	}
	writeToPoti(intensity)
	lm.BreathState = false
	fmt.Println(intensity)
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

func writeToPoti(i int) {
	write := []byte{0x00, byte(i)}
	read := make([]byte, len(write))

	if err := Modules.c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}

var Modules ModulesStruct

func SetupLight() {
	// Make sure periph is initialized.
	host.Init()
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}
	p, err := spireg.Open("SPI0.0")
	if err != nil {
		log.Fatal(err)
	}
	//TODO defer p.Close()
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}

	// Add one Module
	module1 := Module{22, false, make(chan bool), false}
	module1.init()
	module2 := Module{27, false, make(chan bool), false}
	module2.init()

	// Add Modules to Global Variable
	Modules.c = c
	Modules.Modules = make(map[string]Module)
	Modules.Modules["1"] = module1
	Modules.Modules["2"] = module2

}
