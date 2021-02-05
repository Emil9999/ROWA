package sensor

import (
	piblaster "github.com/ddrager/go-pi-blaster"
)

func BlinkLight() {
	a := []int64{4, 22, 24}
	piblaster.Start(a)

	piblaster.Apply(4, 1)

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

	/*for i := 0; i < 10; i++ {
		for i := 1; i < 100; i += 5 { // increasing brightness
			//pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
			x := float64(i) / 100
			val := `"4=` + fmt.Sprint(x) + `"`
			fmt.Print(val)
			cmd := exec.Command("echo", val, ">", "/dev/pi-blaster")
			stdout, err := cmd.Output()
			fmt.Println(string(stdout))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
		for i := 100; i > 0; i -= 5 { // decreasing brightness
			x := float64(i) / 100
			val := `"4=` + fmt.Sprint(x) + `"`
			fmt.Print(val)

			cmd := exec.Command("echo", val, ">", "/dev/pi-blaster")
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(stdout)
			////pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
			time.Sleep(time.Second)

		}

	}*/
}
