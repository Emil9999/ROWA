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

	defer hx711.Shutdown()

	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}

	hx711.AdjustZero, _ = hx711.ReadDataMedianRaw(11)

	hx711.AdjustScale = 140
	for i := 0; i < 10000; i++ {
		time.Sleep(200 * time.Microsecond)

		previousReadings := []float64{}
		movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
		if err != nil {
			fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
		}
		fmt.Println(movingAvg)

	}

	/*var weight1 float64
	var weight2 float64

	weight1 = 1000
	weight2 = 2000

	hx711.GetAdjustValues(weight1, weight2)*/
	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	//hx711.SetGain(128)

	// make sure to use your values from calibration above
	//hx711.AdjustZero = -185820

	//previousReadings = previousReadings.append(data)*/
	//fmt.Println(data)
	/*previousReadings := []float64{}
	movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
	if err != nil {
		fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
	}*/

	// moving average

}

func TestScale() {
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

	defer hx711.Shutdown()

	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}

	var data int
	for i := 0; i < 10000; i++ {
		time.Sleep(2000 * time.Microsecond)

		data, err = hx711.ReadDataRaw()
		if err != nil {
			fmt.Println("ReadDataRaw error:", err)
			continue
		}

		fmt.Println(data)
	}

}
