package sensor

import "time"

func BlinkLight(pin int64) {
	//TODO put module light pins
	a := []int64{4, 22, 24}
	var b Blaster
	b.StartBlaster(a)

	b.ApplyBlaster(pin, 1)
	time.Sleep(time.Second)
	b.ApplyBlaster(pin, 0)

	for {
		for i := 1; i < 100; i++ { // increasing brightness
			//pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
			//x := float64(i) / 100
			b.ApplyBlaster(pin, float64(i)/100)
			time.Sleep(time.Millisecond * 5)
		}
		for i := 100; i > 0; i-- { // decreasing brightness
			//x := float64(i) / 100
			b.ApplyBlaster(pin, float64(i)/100)
			time.Sleep(time.Millisecond * 5)

		}

	}
}
