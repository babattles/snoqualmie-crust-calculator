package crust

import (
	"github.com/babattles/snoqualmie-crust-calculator/inversion"
	"github.com/babattles/snoqualmie-crust-calculator/models"
	"github.com/babattles/snoqualmie-crust-calculator/sun"
)

// returns an array of bools where the index is true when there likely
// exists a sun crust based on temperature inversions & sun effect
func FindSunCrust(data []models.WeatherStationData) []CrustType {
	res := make([]CrustType, len(data))
	inversionsBelow := inversion.FindInversionsBelow(data)
	sunExposures := sun.FindSunEffect(data)
	for i, layer := range(data) {
		gotSun := sunExposures[i]
		inversionBelow := inversionsBelow[i]

		// We have high confidence of a sun crust at the current layer if:
		// * There was a temperature inversion with the layer below
		// * The layer likely recieved sun exposure
		// * The current temperature is above freezing
		if inversionBelow && 
		gotSun && 
		!layer.BelowFreezing() {
			res[i] = CrustSun
			continue
		}

		// Because our sun exposure estimates can't take into account cloud coverage outside our station data,
		// we don't trust a lack of humidity as a failsafe means of determining cloudbreak
		//
		// Here we return a maybe
		if gotSun && !layer.BelowFreezing() {
			res[i] = CrustSunMaybe
			continue
		}

		res[i] = CrustNone
	}
	return res
}

// returns an array of confidences where the index is true when it is likely
// a melt crust will form based on current temperature
//
// NOTE: assumes the temperature will later return to below freezing at night
func FindMeltCrust(data []models.WeatherStationData) []CrustType {
	res := make([]CrustType, len(data))
	for i, layer := range(data) {
		if !layer.BelowFreezing() {
			res[i] = CrustMelt
			continue
		}

		res[i] = CrustNone
	}
	return res
}