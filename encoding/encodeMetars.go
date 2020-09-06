package encoding

import (
	"fmt"

	"github.com/MWein/METARMap/metar"
)

const WindSpeedMode = 2

var Red = []int{1, 0, 0}
var Green = []int{0, 1, 0}
var Blue = []int{0, 0, 1}

var Yellow = []int{1, 1, 0}
var Magenta = []int{1, 0, 1}
var Cyan = []int{0, 1, 1}

var White = []int{1, 1, 1}
var Off = []int{0, 0, 0}

const FlightCatMode = 0

var FlightCatMap = map[string][]int{
	"VFR":  Green,
	"MVFR": Blue,
	"IFR":  Red,
	"LIFR": Magenta,
	"None": Off,
}

const SkyConditionMode = 1

var SkyConditionMap = map[string][]int{
	"SKC": Off,
	"CLR": Off,
	"FEW": Blue,
	"SCT": Yellow,
	"BKN": Red,
	"OVC": White,
}

func EncodeMetars(metars []metar.Metar, mode int) {
	var shiftRegisterBits []int

	for x := 0; x < len(metars); x++ {
		metar := metars[x]

		if mode == FlightCatMode {
			shiftRegisterBits = append(shiftRegisterBits, FlightCatMap[metar.FlightCategory]...)
		} else if mode == SkyConditionMode {
			shiftRegisterBits = append(shiftRegisterBits, SkyConditionMap[metar.SkyCondition]...)
		} else if mode == WindSpeedMode {
			windSpeed := metar.WindSpeed

			if windSpeed <= 5 {
				shiftRegisterBits = append(shiftRegisterBits, Green...)
			} else if windSpeed <= 15 {
				shiftRegisterBits = append(shiftRegisterBits, Yellow...)
			} else if windSpeed > 15 {
				shiftRegisterBits = append(shiftRegisterBits, Red...)
			}
		}
	}

	fmt.Println(shiftRegisterBits)
}
