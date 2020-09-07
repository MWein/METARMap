package main

import (
	//"fmt"

	"github.com/MWein/METARMap/encoding"
	"github.com/MWein/METARMap/metar"
)

func main() {
	var metars = metar.GetTestMetars()
	var encodedMetar = encoding.EncodeMetars(metars, 0)
	var reversedMetar = encoding.InverseEncodedMetars(encodedMetar)

	encoding.WriteToGPIO(reversedMetar)
}
