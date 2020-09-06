package main

import (
	//"fmt"

	"github.com/MWein/METARMap/metar"
	"github.com/MWein/METARMap/encoding"
)

func main() {
	var metars = metar.GetMetars()

	encoding.EncodeMetars(metars, 0)
}
