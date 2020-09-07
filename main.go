package main

import (
	//"fmt"

	"github.com/MWein/METARMap/metar"
	"github.com/MWein/METARMap/encoding"
)

func main() {
	var metars = metar.GetMetars()
	var encodedMetar = encoding.EncodeMetars(metars, 0)
	var reversedMetar = encoding.InverseEncodedMetars(encodedMetar)

	encoding.WriteToGPIO(reversedMetar)
}
