package sensor

import (
	"log"

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
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}

	for i in range(0x00, 0x1FF, 1){
		fmt.Println(i)
		write_pot(i)
		time.sleep(0.5)
	}

	for i in range(0x1FF, 0x00, -1){
		write_pot(i)
		time.sleep(0.5)
	}
	

	// Write 0x10 to the device, and read a byte right after.
	
	// Use read.
	//fmt.Printf("%v\n", read[1:])
}

func write_pot(i byte) {
	write := []byte{i, 0x00}
	read := make([]byte, len(write))
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}