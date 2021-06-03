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
	Module1 Module
	Module2 Module
	Module3 Module
	Module4 Module
	Module5 Module
	Module6 Module
	c       conn.Conn
}

type Module struct {
	Pin           int
	State         bool
	StopBreathing chan bool
	BreathState   bool
}

var (
	globalIntensity  = 0
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
	var arr = [6]Module{ms.Module1, ms.Module2, ms.Module3, ms.Module4, ms.Module5, ms.Module6}
	fmt.Println("called pin", pin)
	for _, module := range arr {
		/*if module.BreathState {
			module.StopBreathing <- true
		}*/
		if module.Pin != pin {

			//fmt.Println("high", module.Pin)
			p := gpioreg.ByName("GPIO" + strconv.Itoa(module.Pin))
			if err := p.Out(gpio.High); err != nil {
				log.Fatal(err)
			}
		} else {
			//fmt.Println("low", module.Pin)
			p := gpioreg.ByName("GPIO" + strconv.Itoa(module.Pin))
			if err := p.Out(gpio.Low); err != nil {
				log.Fatal(err)
			}
		}

	}
}

func LightSwitch(state bool) {
	/*for _, module := range Modules.Modules {
		if state {
			module.LightOn()
		} else {
			module.LightOff()
			}
	}*/
}

func (lm *Module) LightOn() {
	Modules.SetPinsHigh(lm.Pin)
	Increment()
	//writeToPoti(globalIntensity)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) LightOff() {
	fmt.Println("State before", lm.State)
	Modules.SetPinsHigh(lm.Pin)
	writeToPoti(255)
	lm.State = false
	fmt.Println("State", lm.State)
}

func (lm *Module) BreathOn() {
	fmt.Println("Start breathing")
	fmt.Println("State", lm.State)
	fmt.Println("State", lm.Pin)
	intensityDown := lm.State
	var intensity int
	if lm.State {
		intensity = globalIntensity
	} else {
		intensity = 255
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

			if intensity == globalIntensity || intensity == 255 {
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
	fmt.Println("intensitsity is ", intensity)
	writeToPoti(intensity)
	lm.BreathState = false
	fmt.Println(intensity)
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

func writeToPoti(i int) {
	fmt.Println(byte(i))
	write := []byte{0x00, byte(i)}
	read := make([]byte, len(write))

	if err := Modules.c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}
func Increment() {
	for i=:0, i<255 i++ {
		write := []byte{0x01}
		read := make([]byte, len(write))

		if err := Modules.c.Tx(write, read); err != nil {
			log.Fatal(err)
		}
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
	module1 := Module{16, true, make(chan bool), false}
	module1.init()
	module2 := Module{12, true, make(chan bool), false}
	module2.init()
	module3 := Module{25, true, make(chan bool), false}
	module1.init()
	module4 := Module{24, true, make(chan bool), false}
	module2.init()
	module5 := Module{23, true, make(chan bool), false}
	module1.init()
	module6 := Module{18, true, make(chan bool), false}
	module2.init()

	// Add Modules to Global Variable

	Modules.c = c
	Modules.Module1 = module1
	Modules.Module2 = module2
	Modules.Module3 = module3
	Modules.Module4 = module4
	Modules.Module5 = module5
	Modules.Module6 = module6

}
