package sensor

import "time"

func BlinkLight() {

	a := []int64{4, 22, 24}
	var b Blaster
	b.StartBlaster(a)

	b.ApplyBlaster(4, 1)
	time.Sleep(time.Second)
	b.ApplyBlaster(4, 0)

	/*err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	pin := rpio.Pin(26)
	pin.Output()
	pin.Low()
	pin.High()
	time.Sleep(time.Second * 5)
	pin.Low()
	time.Sleep(time.Second * 2)

	for {
		pin.High()
		time.Sleep(time.Millisecond)
		pin.Low()
		time.Sleep(time.Millisecond)
	}*/
	// the LED will be blinking at 2000Hz
	// (source frequency divided by cycle length => 64000/32 = 2000)

	// five times smoothly fade in and out
	/*f, err := os.OpenFile("/dev/pi-blaster", os.O_RDWR, 064)
	fmt.Println(f)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("4=1"); err != nil {
		log.Println(err)
	}*/

	for i := 0; i < 10; i++ {
		for i := 1; i < 100; i += 5 { // increasing brightness
			//pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
			x := float64(i) / 100
			b.ApplyBlaster(4, x)
			time.Sleep(time.Second)
		}
		for i := 100; i > 0; i -= 5 { // decreasing brightness
			x := float64(i) / 100
			b.ApplyBlaster(4, x)
			time.Sleep(time.Second)

		}

	}
}
