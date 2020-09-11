package ledtest

import (
	"fmt"
	"time"

	"github.com/MWein/METARMap/encoding"
	"github.com/MWein/METARMap/metar"
)

func indivAirportLocationTest(airport string) {
	var bits []bool

	for x := 0; x < len(metar.Airports); x++ {
		if airport == metar.Airports[x] {
			bits = append(bits, true, false, false)
		} else {
			bits = append(bits, false, false, false)
		}
	}

	bits = encoding.InverseEncodedMetars(bits)
	encoding.WriteToGPIO(bits)
}

func colorTest(red bool, green bool, blue bool) {
	var bits []bool

	for x := 0; x < len(metar.Airports); x++ {
		bits = append(bits, red, green, blue)
	}

	bits = encoding.InverseEncodedMetars(bits)
	encoding.WriteToGPIO(bits)
}

func RunTests() {
	fmt.Printf("\033[?25l" + "\n-- Running Tests --\n\n")

	fmt.Println("Airport Location Test")
	for x := 0; x < len(metar.Airports); x++ {
		var airport = metar.Airports[x]

		fmt.Printf("\r\033[K" + "Lighting: " + airport)
		indivAirportLocationTest(airport)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\r\033[K" + "Done")

	fmt.Printf("\n\nColor Test\n")

	fmt.Printf("\r\033[K" + "Red")
	colorTest(true, false, false)
	time.Sleep(1 * time.Second)

	fmt.Printf("\r\033[K" + "Green")
	colorTest(false, true, false)
	time.Sleep(1 * time.Second)

	fmt.Printf("\r\033[K" + "Blue")
	colorTest(false, true, false)
	time.Sleep(1 * time.Second)

	fmt.Printf("\r\033[K" + "Magenta")
	colorTest(true, false, true)
	time.Sleep(1 * time.Second)

	fmt.Printf("\r\033[K" + "Done" + "\033[?25h")
}
