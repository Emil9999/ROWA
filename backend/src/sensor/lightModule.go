package sensor

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

type ModulesStruct struct {
	Module1 Module
	Module2 Module
	Module3 Module
	Module4 Module
	Module5 Module
	Module6 Module
}

type Module struct {
	Pin           int64
	State         bool
	StopBreathing chan bool
	BreathState   bool
}

var (
	globalIntensity  = 0.5
	breathingSpeedMs = 40
	b                Blaster
)

func (lm *Module) init() {

	/*p := gpioreg.ByName("GPIO" + strconv.Itoa(lm.Pin))
	if p == nil {
		log.Fatal("Failed to find GPIO" + strconv.Itoa(lm.Pin))
	}
	if err := p.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}*/

}

func (ms *ModulesStruct) SetPinsHigh(pin int64) {
	/*var arr = [6]Module{ms.Module1, ms.Module2, ms.Module3, ms.Module4, ms.Module5, ms.Module6}
	fmt.Println("called pin", pin)
	for _, module := range arr {
		if module.BreathState {
			module.StopBreathing <- true
		}
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

	}*/
}

func (lm *Module) LightOn() {
	b.ApplyBlaster(lm.Pin, globalIntensity)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) LightOff() {
	fmt.Println("State before", lm.State)
	b.ApplyBlaster(lm.Pin, 0)
	lm.State = false
}

func (lm *Module) BreathOn() {
	fmt.Println("Start breathing")
	fmt.Println("State", lm.State)
	fmt.Println("State", lm.Pin)
	intensityDown := lm.State
	var intensity float64
	if lm.State {
		intensity = globalIntensity
	} else {
		intensity = 12
	}
	//Modules.SetPinsHigh(lm.Pin)

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

			if intensity == 12 || intensity == globalIntensity {
				intensityDown = !intensityDown
			}
			b.ApplyBlaster(lm.Pin, intensity/100)
			fmt.Println("writing: " + intensity)
			time.Sleep(time.Duration(breathingSpeedMs) * time.Millisecond)
		}
	}
	lm.BreathState = true
}

func (lm *Module) BreathOff() {
	fmt.Println("Stop breathing")
	fmt.Println("State", lm.State)
	lm.StopBreathing <- true
	//Modules.SetPinsHigh(lm.Pin)

	var intensity float64
	if lm.State {
		intensity = 12
	} else {
		intensity = globalIntensity
	}
	fmt.Println("intensitsity is ", intensity)
	b.ApplyBlaster(lm.Pin, intensity/100)
	lm.BreathState = false
	fmt.Println(intensity)
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

func writeToPoti(i int) {
	/*	fmt.Println(byte(i))
		write := []byte{0x00, byte(i)}
		read := make([]byte, len(write))

		if err := Modules.c.Tx(write, read); err != nil {
			log.Fatal(err)
		}*/

}

var Modules ModulesStruct

func SetupLight() {
	a := []int64{17, 22, 24}
	b.StartBlaster(a)
	// Make sure periph is initialized.
	/*host.Init()
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
	}*/

	// Add one Module
	module1 := Module{16, true, make(chan bool), false}
	module1.init()
	module2 := Module{12, true, make(chan bool), false}
	module2.init()
	module3 := Module{25, true, make(chan bool), false}
	module3.init()
	module4 := Module{24, true, make(chan bool), false}
	module4.init()
	module5 := Module{23, true, make(chan bool), false}
	module5.init()
	module6 := Module{18, true, make(chan bool), false}
	module6.init()

	// Add Modules to Global Variable

	//Modules.c = c
	Modules.Module1 = module1
	Modules.Module2 = module2
	Modules.Module3 = module3
	Modules.Module4 = module4
	Modules.Module5 = module5
	Modules.Module6 = module6

}

func (b *Blaster) SetBrightness(pin int64, brightness float64) {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	pin1 := rpio.Pin(pin)
	pin1.Output()
	pin1.High()

	b.ApplyBlaster(pin, brightness)
}
func (b *Blaster) BlinkLight(pin int64, toggle bool) {
	/*err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	pin1 := rpio.Pin(pin)
	pin1.Output()
	pin1.High()*/

	//TODO put module light pins

	b.ApplyBlaster(pin, 0.7)
	time.Sleep(time.Second * 2)
	if toggle {
		for {
			for i := 12; i < 70; i++ { // increasing brightness
				b.ApplyBlaster(pin, float64(i)/100)
				time.Sleep(time.Millisecond * 20)
			}
			for i := 70; i > 11; i-- { // decreasing brightness
				b.ApplyBlaster(pin, float64(i)/100)
				time.Sleep(time.Millisecond * 20)

			}
		}

	} else {

	}
}
