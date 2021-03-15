package sensor

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)
var (
	lights = [27,22]
	c conn.Conn
)

func Spiinit() {
	// Make sure periph is initialized.
	host.Init()
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	p, err := spireg.Open("SPI0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	c, err = p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}
}

func BreatheLight(pin int){
	p:= gpioreg.ByName("GPIO" + strconv.Itoa(pin))
		if p == nil {
			log.Fatal("Failed to find GPIO" + strconv.Itoa(pin) )
		}
	if err := p.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < 256; i++ {
			wroteToPoti(i)
			time.Sleep(time.Millisecond * 5)
		}

		for i := 255; i > -1; i-- {
			writeToPoti(i)
			time.Sleep(time.Millisecond * 5)
		}
	}
}

func wroteToPoti(i int) {
	write := []byte{0x00, byte(i)}
	read := make([]byte, len(write))
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}

func initLightPins(lights []int) {
	//Set all light pins to HIGH
	for i in lights {
		p:= gpioreg.ByName("GPIO" + strconv.Itoa(lights[i]))
		if p == nil {
			log.Fatal("Failed to find GPIO" + strconv.Itoa(lights[i]) )
		}
		// Set the pin as output High.
		if err := p.Out(gpio.High); err != nil {
		log.Fatal(err)
		}	

	}
	
}