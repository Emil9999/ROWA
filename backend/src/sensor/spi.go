package sensor

import (
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
)

func Spiinit() {
	// Make sure periph is initialized.
	// TODO: Use host.Init(). It is not used in this example to prevent circular
	// go package import.
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use spireg SPI port registry to find the first available SPI bus.
	p, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}

	/*for i in range(0x00, 0x1FF, 1){
		write_pot(i)
		time.sleep(.005)
	}

	for i in range(0x1FF, 0x00, -1):
	write_pot(i)
	time.sleep(.005)*/

	// Write 0x10 to the device, and read a byte right after.
	write := []byte{0x10, 0x00}
	if err := c.Tx(write); err != nil {
		log.Fatal(err)
	}
	// Use read.
	//fmt.Printf("%v\n", read[1:])
}
