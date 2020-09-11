package main

import (
	"fmt"

	"time"

	"github.com/MWein/METARMap/encoding"
	"github.com/MWein/METARMap/ledtest"
	"github.com/MWein/METARMap/metar"
)

func main() {
	ledtest.RunTests()

	for {
		var metars = metar.GetMetars()

		if len(metars) != 0 {
			var encodedMetar = encoding.EncodeMetars(metars, 0)
			var reversedMetar = encoding.InverseEncodedMetars(encodedMetar)
			encoding.WriteToGPIO(reversedMetar)
		} else {
			fmt.Println("Error. Trying again in 5 minutes")
		}

		// Wait for 10 minutes
		time.Sleep(10 * time.Minute)
	}
}
