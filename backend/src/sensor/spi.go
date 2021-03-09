package sensor

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func Spiinit() {
	// Make sure periph is initialized.
	// TODO: Use host.Init(). It is not used in this example to prevent circular
	// go package import.
	host.Init()
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
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 16)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 258; i++ {
		fmt.Println(i)
		write_pot(i, c)
		time.Sleep(time.Second)
	}

	for i := 257; i > 0; i-- {
		fmt.Println(i, c)
		write_pot(i)
		time.Sleep(time.Second)
	}

	// Write 0x10 to the device, and read a byte right after.

	// Use read.
	//fmt.Printf("%v\n", read[1:])
}

func write_pot(i int, c Conn) {

	write := []byte{i, 0x00}
	read := make([]byte, len(write))
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}
