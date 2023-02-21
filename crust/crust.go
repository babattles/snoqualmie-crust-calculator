package crust

import (
	"github.com/babattles/snoqualmie-crust-calculator/inversion"
	"github.com/babattles/snoqualmie-crust-calculator/models"
	"github.com/babattles/snoqualmie-crust-calculator/sun"
)

type CrustConfidence string

const (
	CrustYes CrustConfidence = "yes"
	CrustNo CrustConfidence = "no"
	CrustMaybe CrustConfidence = "maybe"
)

// returns an array of bools where the index is true when there likely
// exists a sun crust based on temperature inversions & sun effect
func FindSunCrust(data []models.WeatherStationData) []CrustConfidence {
	res := make([]CrustConfidence, len(data))
	inversionsBelow := inversion.FindInversionsBelow(data)
	sunExposures := sun.FindSunEffect(data)
	for i, layer := range(data) {
		gotSun := sunExposures[i]
		inversionBelow := inversionsBelow[i]

		// if there was a temperature inversion detected below
		// AND this layer might have recieved sun exposure
		// AND the temperature is above freezing
		// there is very likely a sun crust
		// NOTE: this assumption is likely to upset many people
		if inversionBelow && 
		gotSun && 
		!layer.BelowFreezing() {
			res[i] = CrustYes
			continue
		}

		// because our sun exposure estimate isn't perfect, we provide a maybe if we can't 
		// guess more precisely because we detected an inversion
		// so if the layer might have received sun exposure 
		// AND is above freezing, it might have a crust
		// NOTE: this check assumes mid-winter conditions that will return this layer to freezing at some point
		// YMMV
		if gotSun && !layer.BelowFreezing() {
			res[i] = CrustMaybe
			continue
		}

		res[i] = CrustNo
	}
	return res
}