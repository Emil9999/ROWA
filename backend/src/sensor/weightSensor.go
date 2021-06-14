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

	hx711.AdjustZero, _ = hx711.ReadDataMedianRaw(11)

	hx711.AdjustScale = 140

	return hx711, err
}
func ReadWeight(hx711 *hx711.Hx711) float64 {

	defer hx711.Shutdown()
	time.Sleep(200 * time.Microsecond)
	previousReadings := []float64{}
	movingAvg, err := hx711.ReadDataMedianThenMovingAvgs(11, 8, &previousReadings)
	if err != nil {
		fmt.Println("ReadDataMedianThenMovingAvgs error:", err)
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

	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)

	// make sure to use your values from calibration above
	hx711.AdjustZero = -123
	hx711.AdjustScale = 456

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
