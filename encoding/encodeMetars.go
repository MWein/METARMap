package encoding

import (
	"fmt"

	"github.com/MWein/METARMap/metar"
)

const FlightCatMode = 0
const SkyConditionMode = 1
const WindSpeedMode = 2

var red = []bool{true, false, false}
var green = []bool{false, true, false}
var blue = []bool{false, false, true}

var yellow = []bool{true, true, false}
var magenta = []bool{true, false, true}
var cyan = []bool{false, true, true}

var white = []bool{true, true, true}
var off = []bool{false, false, false}

var flightCatMap = map[string][]bool{
	"VFR":  green,
	"MVFR": blue,
	"IFR":  red,
	"LIFR": magenta,
	"None": off,
}

var skyConditionMap = map[string][]bool{
	"SKC": off,
	"CLR": off,
	"FEW": blue,
	"SCT": yellow,
	"BKN": red,
	"OVC": white,
	"None": off,
}

func EncodeMetars(metars []metar.Metar, mode int) {
	var shiftRegisterBits []bool

	for x := 0; x < len(metars); x++ {
		metar := metars[x]

		if mode == FlightCatMode {
			shiftRegisterBits = append(shiftRegisterBits, flightCatMap[metar.FlightCategory]...)
		} else if mode == SkyConditionMode {
			shiftRegisterBits = append(shiftRegisterBits, skyConditionMap[metar.SkyCondition]...)
		} else if mode == WindSpeedMode {
			windSpeed := metar.WindSpeed

			if windSpeed <= 5 {
				shiftRegisterBits = append(shiftRegisterBits, green...)
			} else if windSpeed <= 15 {
				shiftRegisterBits = append(shiftRegisterBits, yellow...)
			} else if windSpeed > 15 {
				shiftRegisterBits = append(shiftRegisterBits, red...)
			} else {
				shiftRegisterBits = append(shiftRegisterBits, off...)
			}
		}
	}

	fmt.Println(shiftRegisterBits)
}
