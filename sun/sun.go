package sun

import "github.com/babattles/snoqualmie-crust-calculator/models"

func FindSunEffect(data []models.WeatherStationData) []bool {
	// sort first for peace of mind
	models.SortByElevation(data)

	res := make([]bool, len(data))
	for i := len(data)-1; i >= 0; i-- {
		sunOut := data[i].CloudBreak()
		if sunOut {
			res[i] = true
		} else {
			// sun isn't out, everything below won't experience sun
			res[i] = false
			for j := i-1; j >= 0; j-- {
				res[j] = false
			}
			return res
		}
	 }

	return res
}