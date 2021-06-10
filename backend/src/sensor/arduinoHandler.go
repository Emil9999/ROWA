package sensor

import (
	"log"
	"time"

	"github.com/tarm/serial"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"

	"periph.io/x/periph/experimental/devices/pca9685"
	host "periph.io/x/periph/host"
)

var ch = make(chan string, 100)

func WriteToCh(input string) {
	ch <- input
}

func ArduinoLoop(s *serial.Port) {

	log.Println("Starting loop")
	_, err := s.Write([]byte("80"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("81"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("80"))
	time.Sleep(time.Second * 1)

	for {

		for v := range ch {
			log.Println("Writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)

		}
	}
}

func PwmTest() {
	/*	options := &pca9685.Options{"pca0", 800.0, 25000000.0}

		// Create new connection to i2c-bus on 1 line with address 0x40.
		// Use i2cdetect utility to find device address over the i2c-bus
		i2c, err := i2c.New(pca9685.Address, 1)
		if err != nil {
			log.Fatal(err)
		}

		pca0, err := pca9685.New(i2c, options)
		if err != nil {
			log.Fatal(err)
		}
		//pca0.SetFreq(800.0)

		// Sets a single PWM channel 0
		pca0.SetChannel(0, 0, 4094)

		// Servo on channel 0
		light1 := pca0.ServoNew(0, nil)
		light1.Fraction(1)

		//fmt.Println(pca0.GetFreq())*/
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	pca, err := pca9685.NewI2C(bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}

	if err := pca.SetPwmFreq(800 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := pca.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}

	if err := pca.SetPwm(15, 0, 4095); err != nil {
		log.Fatal(err)
	}

}
