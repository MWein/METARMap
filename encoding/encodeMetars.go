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
	"SKC":  off,
	"CLR":  off,
	"FEW":  blue,
	"SCT":  yellow,
	"BKN":  red,
	"OVC":  white,
	"None": off,
}

// This function is used if you're an idiot like me and bought multicolor LEDs where the anode and cathode is reversed
// IE, red should be 100... but the damn LED lights red for 011 because I bought the wrong kind
func InverseEncodedMetars(oldBits []bool) []bool {
	var newBits []bool
	for x := 0; x < len(oldBits); x++ {
		newBits = append(newBits, !oldBits[x])
	}
	return newBits
}

// TODO Change to reference the actual airports list
// Metar objects should be in a map with airport code as the key, metar as the value

// TODO Maybe get rid of sky condition mode. Unless I can find a way to control brightness so it can all show as white
func EncodeMetars(metars []metar.Metar, mode int) []bool {
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
	fmt.Println(InverseEncodedMetars(shiftRegisterBits))

	return shiftRegisterBits
}
