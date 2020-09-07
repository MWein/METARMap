package encoding

import (
	"github.com/stianeikeland/go-rpio/v4"
)

const DATALINE = 23
const OCLK = 18
const ICLK = 15

func WriteToGPIO(bits []bool) {
	// Open GPIO access
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	// Initialize Pins
	dlPin := rpio.Pin(DATALINE)
	oClkPin := rpio.Pin(OCLK)
	iClkPin := rpio.Pin(ICLK)
	dlPin.Output()
	oClkPin.Output()
	iClkPin.Output()

	// Push the bits to the shift registers
	for x := 0; x < len(bits); x++ {
		if bits[x] {
			dlPin.High()
			iClkPin.High()
			iClkPin.Low()
			dlPin.Low()
		} else {
			iClkPin.High()
			iClkPin.Low()
		}
	}

	// Tell the shift registers to display
	oClkPin.High()

	// TODO Remove comment out when I replace the bad register
	//oClkPin.Low()
}
