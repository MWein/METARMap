package encoding

import (
	"fmt"

	"github.com/MWein/METARMap/metar"
)

const FlightCatMode = 0
const SkyConditionMode = 1
const WindSpeedMode = 2

var red = []int{1, 0, 0}
var green = []int{0, 1, 0}
var blue = []int{0, 0, 1}

var yellow = []int{1, 1, 0}
var magenta = []int{1, 0, 1}
var cyan = []int{0, 1, 1}

var white = []int{1, 1, 1}
var off = []int{0, 0, 0}

var flightCatMap = map[string][]int{
	"VFR":  green,
	"MVFR": blue,
	"IFR":  red,
	"LIFR": magenta,
	"None": off,
}

var skyConditionMap = map[string][]int{
	"SKC": off,
	"CLR": off,
	"FEW": blue,
	"SCT": yellow,
	"BKN": red,
	"OVC": white,
}

func EncodeMetars(metars []metar.Metar, mode int) {
	var shiftRegisterBits []int

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
			}
		}
	}

	fmt.Println(shiftRegisterBits)
}
