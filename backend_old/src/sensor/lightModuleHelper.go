//Helper Funcs
package sensor

import (
	"fmt"
	"time"
)

func LightOnModule(ModuleNum int) {
	if DetectRpi() {
	switch ModuleNum {
	case 1:
		Modules.Module1.LightOn()
	case 2:
		Modules.Module2.LightOn()
	case 3:
		Modules.Module3.LightOn()
	case 4:
		Modules.Module4.LightOn()
	case 5:
		Modules.Module5.LightOn()
	case 6:
		Modules.Module6.LightOn()
	default:
		fmt.Println("Module not implemented")
	}
} else {
	fmt.Println("No Rpi Present")
}

}

func LightOffModule(ModuleNum int) {
	if DetectRpi() {
	switch ModuleNum {
	case 1:
		Modules.Module1.LightOff()
	case 2:
		Modules.Module2.LightOff()
	case 3:
		Modules.Module3.LightOff()
	case 4:
		Modules.Module4.LightOff()
	case 5:
		Modules.Module5.LightOff()
	case 6:
		Modules.Module6.LightOff()
	default:
		fmt.Println("Module not implemented")
	}
	} else {
		fmt.Println("No Rpi Present")
	}

}

func BreathOnModule(ModuleNum int) {
	if DetectRpi() {
	switch ModuleNum {
	case 1:
		go Modules.Module1.BreathOn()
	case 2:
		go Modules.Module2.BreathOn()
	case 3:
		go Modules.Module3.BreathOn()
	case 4:
		go Modules.Module4.BreathOn()
	case 5:
		go Modules.Module5.BreathOn()
	case 6:
		go Modules.Module6.BreathOn()
	default:
		fmt.Println("Module not implemented")
	}
	} else {
		fmt.Println("No Rpi Present")
	}

}

func BreathOffModule(ModuleNum int) {
	if DetectRpi() {
	switch ModuleNum {
	case 1:
		Modules.Module1.BreathOff()
	case 2:
		Modules.Module2.BreathOff()
	case 3:
		Modules.Module3.BreathOff()
	case 4:
		Modules.Module4.BreathOff()
	case 5:
		Modules.Module5.BreathOff()
	case 6:
		Modules.Module6.BreathOff()
	default:
		fmt.Println("Module not implemented")
	}
	} else {
		fmt.Println("No Rpi Present")
	}

}

func LightAllOn() {
	if DetectRpi() {
	for i := 1; i < 7; i++ {
		LightOnModule(i)
		time.Sleep(time.Millisecond * 100)
	}
	} else {
		fmt.Println("No Rpi Present")
	}
}

func LightAllOff() {
	if DetectRpi() {
	for i := 1; i < 7; i++ {
		LightOffModule(i)
		time.Sleep(time.Millisecond * 100)

	}
	} else {
		fmt.Println("No Rpi Present")
	}
}
