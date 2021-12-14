package sensor

import (
	"fmt"
	"time"
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
	globalIntensity  = 0.6
	breathingSpeedMs = 30
	b                Blaster
)

func (lm *Module) LightOn() {
	b.ApplyBlaster(lm.Pin, globalIntensity)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) LightOff() {
	//fmt.Println("State before", lm.State)
	b.ApplyBlaster(lm.Pin, 0)
	lm.State = false
}

func (lm *Module) BreathOn() {

	fmt.Println("Start breathing")
	fmt.Println(len(lm.StopBreathing))
	//fmt.Println("State", lm.State)
	intensityDown := lm.State
	var intensity float64
	if lm.State {
		intensity = globalIntensity * 100
	} else {
		intensity = 5
	}
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
			if intensityDown {
				intensity--
			} else {
				intensity++
			}

			if intensity == 5 || intensity == globalIntensity*100 {
				intensityDown = !intensityDown
			}
			b.ApplyBlaster(lm.Pin, intensity/100)
			time.Sleep(time.Duration(breathingSpeedMs) * time.Millisecond)
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
	
	var intensity float64
	if lm.State {
		intensity = globalIntensity
	} else {
		intensity = 0
	}
	b.ApplyBlaster(lm.Pin, intensity)
	lm.BreathState = false
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

var Modules ModulesStruct

func SetupLight() {
	a := []int64{16, 12, 25, 23, 18, 24}
	b.StartBlaster(a)

	// Add one Module
	module1 := Module{16, true, make(chan bool, 3), false}
	module2 := Module{12, true, make(chan bool, 3), false}
	module3 := Module{25, true, make(chan bool, 3), false}
	module4 := Module{24, true, make(chan bool, 3), false}
	module5 := Module{23, true, make(chan bool, 3), false}
	module6 := Module{18, true, make(chan bool, 3), false}

	// Add Modules to Global Variable

	Modules.Module1 = module1
	Modules.Module2 = module2
	Modules.Module3 = module3
	Modules.Module4 = module4
	Modules.Module5 = module5
	Modules.Module6 = module6

}
