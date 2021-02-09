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
	weight2 = 2000

	hx711.GetAdjustValues(weight1, weight2)*/
	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	//hx711.SetGain(128)

	// make sure to use your values from calibration above
	//hx711.AdjustZero = -185820
	/*hx711.AdjustZero, _ = hx711.ReadDataMedianRaw(11)

	hx711.AdjustScale = 1050
	var data float64
	previousReadings := []float64{}

	for i := 0; i < 12; i++ {
		time.Sleep(200 * time.Microsecond)

		data, err = hx711.ReadDataMedian(11)
		if err != nil {
			fmt.Println("ReadDataMedian error:", err)
			continue
		}
		previousReadings = previousReadings.append(data)
		//fmt.Println(data)
	}
	movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
	if err != nil {
		fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
	}

	// moving average
	fmt.Println(movingAvg)*/

	hx711.AdjustZero, _ = hx711.ReadDataMedianRaw(11)

	hx711.AdjustScale = 1050
	var movingAvg float64
	stop := false
	stopped = make(chan struct{})
	go hx711.BackgroundReadMovingAvgs(11, 8, &movingAvg, &stop, stopped)

	// wait for data
	time.sleep(time.Second)

	// moving average
	fmt.Println(movingAvg)

	// when done set stop to true to quit BackgroundReadMovingAvgs
	stop = true

	// wait for BackgroundReadMovingAvgs to stop
	<-stopped
}
