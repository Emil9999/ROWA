package sensor

import (
	"fmt"
	"log"
	"os/exec"
)

func BlinkLight() {
	/*err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	pin := rpio.Pin(13)*/
	/*pin.Output()
	pin.High()
	time.Sleep(time.Second * 10)
	pin.Low()
	pin.Pwm()
	pin.Freq(64000)
	pin.DutyCycle(10, 75)*/
	// the LED will be blinking at 2000Hz
	// (source frequency divided by cycle length => 64000/32 = 2000)

	// five times smoothly fade in and out
	for i := 0; i < 10; i++ {
		for i := 1; i < 100; i += 5 { // increasing brightness
			//pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
			x := float64(i) / 100
			val := `"4=` + fmt.Sprint(x) + `"`
			fmt.Print(val)
			cmd := exec.Command("echo", val, ">", "/dev/pi-blaster")
			_, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
		}
		for i := 100; i > 0; i -= 5 { // decreasing brightness
			x := float64(i) / 100
			val := `"4=` + fmt.Sprint(x) + `"`
			fmt.Print(val)

			cmd := exec.Command("echo", val, ">", "/dev/pi-blaster")
			_, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			////pin.DutyCycle(i, 32)
			//time.Sleep(time.Second / 32)
		}
	}
}
