package models

import "sort"

type WeatherStationData struct {
	ElevationFt int
	TemperatureF int
	RelativeHumidityPercent int
}

func SortByElevation(data []WeatherStationData) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].LowerThan(data[j])
	})
}

func (w WeatherStationData) LowerThan(band WeatherStationData) bool {
	return w.ElevationFt < band.ElevationFt
}

func (w WeatherStationData) CloudBreak() bool {
	return w.RelativeHumidityPercent < CloudBreakHumidityThreshold
}

func (w WeatherStationData) BelowFreezing() bool {
	return w.TemperatureF < FreezingTempF
}