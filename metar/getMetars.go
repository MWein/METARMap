package metar

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SkyConditionXML struct {
	XMLName      xml.Name `xml:"sky_condition"`
	SkyCondition string   `xml:"sky_cover,attr"`
}

type MetarXML struct {
	XMLName      xml.Name        `xml:"METAR"`
	StationID    string          `xml:"station_id"`
	WindSpeed    int             `xml:"wind_speed_kt"`
	WindGust     int             `xml:"wind_gust_kt"`
	FlightCat    string          `xml:"flight_category"`
	SkyCondition SkyConditionXML `xml:"sky_condition"`
}

type DataXML struct {
	XMLName xml.Name   `xml:"data"`
	Metar   []MetarXML `xml:"METAR"`
}

type ResponseXML struct {
	XMLName xml.Name `xml:"response"`
	Data    DataXML  `xml:"data"`
}

type Metar struct {
	StationID      string
	FlightCategory string
	SkyCondition   string
	WindSpeed      int
}

func GetTestMetars() []Metar {
	var metars []Metar

	for x := 0; x < 5; x++ {
		var metar Metar
		metars = append(metars, metar)
	}

	metars[0].FlightCategory = "VFR"
	metars[1].FlightCategory = "MVFR"
	metars[2].FlightCategory = "IFR"
	metars[3].FlightCategory = "None"
	metars[4].FlightCategory = "LIFR"

	return metars
}

func GetMetars() []Metar {
	// Build to URL from the Airports list
	urlEncodedAirports := strings.Join(Airports, "%20")
	url := "https://www.aviationweather.gov/adds/dataserver_current/httpparam?datasource=metars&requestType=retrieve&format=xml&mostRecentForEachStation=constraint&hoursBeforeNow=1&stationString=" + urlEncodedAirports

	// Request
	resp, err := http.Get(url)
	if err != nil {
		// TODO Handle
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO Handle
	}

	// Read the XML
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO Handle
	}

	// Unmarshal the XML
	var result ResponseXML
	xml.Unmarshal(data, &result)

	// Convert into a more digestible format
	// Also sorts the Metars in the proper order for later encoding
	var metars []Metar
	for x := 0; x < len(Airports); x++ {
		var airport = Airports[x]
		var metar Metar

		// Find the Metar mathing the airport
		var metarXML MetarXML
		for y := 0; y < len(result.Data.Metar); y++ {
			if result.Data.Metar[y].StationID == airport {
				metarXML = result.Data.Metar[y]
				break
			}
		}

		// Create METAR object
		metar.StationID = metarXML.StationID
		metar.FlightCategory = metarXML.FlightCat
		metar.SkyCondition = metarXML.SkyCondition.SkyCondition

		// Take the greater of WindGust or WindSpeed
		if metarXML.WindGust > metarXML.WindSpeed {
			metar.WindSpeed = metarXML.WindGust
		} else {
			metar.WindSpeed = metarXML.WindSpeed
		}

		// Data could be missing
		if metar.FlightCategory == "" {
			metar.FlightCategory = "None"
		}
		if metar.SkyCondition == "" {
			metar.SkyCondition = "None"
		}

		metars = append(metars, metar)
	}

	fmt.Println(len(metars))
	fmt.Println(metars)

	return metars
}
