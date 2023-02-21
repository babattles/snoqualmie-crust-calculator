package inversion

import (
	"errors"

	"github.com/babattles/snoqualmie-crust-calculator/models"
)

var (
	ErrImproperElevationOrdering = errors.New("improper elevation ordering")
)

type InversionData struct {
	LowerElevationFt int
	HigherElevationFt int
	InversionPresent bool
}

// for each elevation band, return if there was an inversion between it and the above elevation band
// (the uppermost elevation band will always be false)
func FindInversionsAbove(data []models.WeatherStationData) []bool {
	// sort first for peace of mind
	models.SortByElevation(data)

	res := make([]bool, len(data))
	for i, layer := range(data) {
		// uppermost layer
		if i == len(data) - 1 {
			res[i] = false
			return res
		}

		res[i] = temperatureInversionExists(layer, data[i+1])
	}

	return res
}

// for each elevation band, return if there was an inversion between it and the elevation band below
// (the lowest elevation band will always be false)
func FindInversionsBelow(data []models.WeatherStationData) []bool {
	// sort first for peace of mind
	models.SortByElevation(data)

	res := make([]bool, len(data))
	for i := len(data)-1; i >= 0; i-- {
		// lowest layer
		if i == 0 {
			res[i] = false
			return res
		}

		res[i] = temperatureInversionExists(data[i-1], data[i])
	}

	return res
}

// calculates if there was a temperature inversion between two elevation bands
func temperatureInversionExists(
	lowerBand models.WeatherStationData, higherBand models.WeatherStationData,
) bool {
	return higherBand.TemperatureF > lowerBand.TemperatureF
}