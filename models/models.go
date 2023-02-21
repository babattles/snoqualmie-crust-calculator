package models

import "sort"

const (
	CloudBreakHumidityThreshold int = 70
	FreezingTempF int = 32
)

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

func (e WeatherStationData) LowerThan(band WeatherStationData) bool {
	return e.ElevationFt < band.ElevationFt
}

func (e WeatherStationData) CloudBreak() bool {
	return e.RelativeHumidityPercent < CloudBreakHumidityThreshold
}

func (e WeatherStationData) BelowFreezing() bool {
	return e.TemperatureF < FreezingTempF
}

