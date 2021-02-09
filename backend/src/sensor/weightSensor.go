package sensor

import (
	"fmt"
	"time"

	"github.com/MichaelS11/go-hx711"
)

func ReadWeight() {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}
	/*var weight1 float64
	var weight2 float64

	weight1 = 1000
	weight2 = 2000*/

	//hx711.GetAdjustValues(weight1, weight2)
	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	//hx711.SetGain(128)

	// make sure to use your values from calibration above
	//hx711.AdjustZero = -185820
	hx711.AdjustZero, _ = hx711.ReadDataMedianRaw(11)

	hx711.AdjustScale = 1000
	var data float64
	for i := 0; i < 10000; i++ {
		time.Sleep(200 * time.Microsecond)

		data, err = hx711.ReadDataMedian(11)
		if err != nil {
			fmt.Println("ReadDataMedian error:", err)
			continue
		}

		fmt.Println(data)
	}
}
