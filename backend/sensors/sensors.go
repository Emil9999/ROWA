package sensors

import (
	"fmt"

	_ "../settings"
	"github.com/stianeikeland/go-rpio"
)

func measureTemp(pin_number int) (temp float32) {
	fmt.Println("Measure Temperature")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(pin_number)
	pin.Input()
	pin.PullUp() // TODO Test it! no idea if pullUp or pullDown, and if this is correct

	voltage := float32(pin.Read())
	temp = (5.0 * voltage * 1000.0) / (1024.0 * 10.0)
	return
}

func measureLight(pin_number int) (lux float32) {
	fmt.Println("Measure Temperature")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(pin_number)
	pin.Input()
	pin.PullUp() // TODO Test it! no idea if pullUp or pullDown, and if this is correct
	lux = float32(pin.Read())
	return
}

//code example from internet
func link() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(21)
	pin.Output()
	pin.Toggle()

	/*for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}*/
}
