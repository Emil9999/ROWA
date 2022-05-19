package sensor

import (
	"fmt"
	"time"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"strconv"
	"github.com/labstack/gommon/log"

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
	breathingSpeedS = 2
)

func (lm *Module) LightOn() {
	pinHigh(lm.Pin)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) LightOff() {
	//fmt.Println("State before", lm.State)
	pinLow(lm.Pin)
	lm.State = false
}

func (lm *Module) BreathOn() {

	fmt.Println("Start breathing")
	fmt.Println(len(lm.StopBreathing))
	//fmt.Println("State", lm.State)
	
	//Emptying channel
	for len(lm.StopBreathing) > 0 {
		<-lm.StopBreathing
	}
	
	for {
		select {
		case <-lm.StopBreathing:
			fmt.Println("Stop infinite loop")
			return
		default:
			pinHigh(lm.Pin)
			time.Sleep(time.Duration(breathingSpeedS + 3) * time.Second)
			pinLow(lm.Pin)
			time.Sleep(time.Duration(breathingSpeedS - 1 ) * time.Second)

		}
	}
	lm.BreathState = true
}

func (lm *Module) BreathOff() {
	
	fmt.Println("State", lm.State)
	if len(lm.StopBreathing) == 0 {
		fmt.Println("Stop breathing")
		lm.StopBreathing <- true
	}
	
	
	if lm.State {
		fmt.Println("pin high")
		lm.LightOn()
		pinHigh(lm.Pin)
	} else {
		fmt.Println("pin low")
		lm.LightOn()
	}
	lm.LightOn()
	lm.BreathState = false
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

var Modules ModulesStruct

func SetupLight() {
	

	// Add one Module
	module1 := Module{16, false, make(chan bool, 3), false}
	module2 := Module{12, false, make(chan bool, 3), false}
	module3 := Module{25, false, make(chan bool, 3), false}
	module4 := Module{24, false, make(chan bool, 3), false}
	module5 := Module{23, false, make(chan bool, 3), false}
	module6 := Module{18, false, make(chan bool, 3), false}

	// Add Modules to Global Variable

	Modules.Module1 = module1
	Modules.Module2 = module2
	Modules.Module3 = module3
	Modules.Module4 = module4
	Modules.Module5 = module5
	Modules.Module6 = module6

}


func pinHigh(pin int64)error{
	p := gpioreg.ByName("GPIO" + strconv.FormatInt(pin, 10))
	if err := p.Out(gpio.Low); err != nil {
		log.Error(err)
		return err
}
return nil
}


func pinLow(pin int64)error{
	p := gpioreg.ByName("GPIO" + strconv.FormatInt(pin, 10))
	if err := p.Out(gpio.High); err != nil {
		log.Error(err)
		return err
}
return nil
}

