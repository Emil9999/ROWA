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

func Spiinit() {
	// Make sure periph is initialized.
	// TODO: Use host.Init(). It is not used in this example to prevent circular
	// go package import.
	host.Init()
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}
	//var o spireg.Opener
	//spireg.Register("GPIO", ["SPI2"], 25, o)
	fmt.Println(spireg.All())
	// Use spireg SPI port registry to find the first available SPI bus.
	//p2, err := bitbang.NewSPI(11, 10, 9, 25)
	/*if err != nil {
		log.Fatal(err)
	}*/
	p1 := gpioreg.ByName("GPIO27")
	if p1 == nil {
		log.Fatal("Failed to find GPIO27")
	}

	// Set the pin as output High.
	if err := p1.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
	p2 := gpioreg.ByName("GPIO22")
	if p2 == nil {
		log.Fatal("Failed to find GPIO22")
	}

	// Set the pin as output High.
	if err := p2.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
	p, err := spireg.Open("SPI0.0")
	fmt.Println(p)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	/*cB, err := p2.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}*/
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}
	if err := p1.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
	if err := p2.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < 256; i++ {
			fmt.Println(i)
			write := []byte{0x00, byte(i)}
			read := []byte{}
			if err := c.Tx(write, read); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 5)
		}

		for i := 255; i > -1; i-- {
			fmt.Println(i)
			write := []byte{0x00, byte(i)}
			read := []byte{}
			if err := c.Tx(write, read); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 5)
		}
	}

	// Write 0x10 to the device, and read a byte right after.

	// Use read.
	//fmt.Printf("%v\n", read[1:])
	/*write := []byte{0x00, 0xC8}
	read := []byte{}
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}*/
	//time.Sleep(time.Millisecond * 100)
}

/*func write_pot(i int, c Conn) {

	write := []byte{i, 0x00}
	read := make([]byte, len(write))
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
}*/
