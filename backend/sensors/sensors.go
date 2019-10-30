package sensors

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

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
