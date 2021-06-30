package sensor

import (
	"fmt"
	"time"

	hx711 "github.com/MichaelS11/go-hx711"
)

func InitScale() (*hx711.Hx711, error) {
	var emptyHx *hx711.Hx711
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return emptyHx, err
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return emptyHx, err
	}
	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return emptyHx, err
	}

	hx711.AdjustZero = -164772
	hx711.AdjustScale = 114

	return hx711, err
}
func ReadWeight(hx711 *hx711.Hx711) float64 {

	defer hx711.Shutdown()
	time.Sleep(200 * time.Microsecond)
	previousReadings := []float64{}
	movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
	if err != nil {
		fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
		movingAvg = -1
	}
	fmt.Println("Weight", movingAvg)
	return movingAvg

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

	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}
	hx711.AdjustZero = -164772

	hx711.AdjustScale = 114
	for i := 0; i < 10000; i++ {
		time.Sleep(200 * time.Millisecond)

		previousReadings := []float64{}
		movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
		if err != nil {
			fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
		}

		// moving average
		fmt.Println(movingAvg)
	}

}

func CalibrateHx() {
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

	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)

	var weight1 float64
	var weight2 float64

	weight1 = 450
	weight2 = 1500

	hx711.GetAdjustValues(weight1, weight2)
}
