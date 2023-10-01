package crust_test

import (
	"testing"

	"github.com/babattles/snoqualmie-crust-calculator/crust"
	"github.com/babattles/snoqualmie-crust-calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestFindSunCrust(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data []models.WeatherStationData
		expected []crust.CrustType
	}{
        {
			name: "no inversions - no clouds - all above freezing - all maybes",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 36,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 500,
					TemperatureF: 34,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 32,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []crust.CrustType{crust.CrustSunMaybe, crust.CrustSunMaybe, crust.CrustSunMaybe},
		},
		{
			name: "no inversions - all clouds - all above freezing - all nos",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 36,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 500,
					TemperatureF: 34,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 32,
					RelativeHumidityPercent: 100,
				},
			},
			expected: []crust.CrustType{crust.CrustNone, crust.CrustNone, crust.CrustNone},
		},
		{
			name: "no inversions - no clouds - all below freezing - all nos",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 28,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 500,
					TemperatureF: 26,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 24,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []crust.CrustType{crust.CrustNone, crust.CrustNone, crust.CrustNone},
		},
		{
			name: "all inversions & all clouds - no crusts",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 30,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 500,
					TemperatureF: 32,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 34,
					RelativeHumidityPercent: 100,
				},
			},
			expected: []crust.CrustType{crust.CrustNone, crust.CrustNone, crust.CrustNone},
		},
		{
			name: "all inversions - no clouds - all below freezing - no crusts",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 12,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 500,
					TemperatureF: 20,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 30,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []crust.CrustType{crust.CrustNone, crust.CrustNone, crust.CrustNone},
		},
		{
			name: "all inversions - no clouds - one below freezing - two crusts",
			data: []models.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 30,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 500,
					TemperatureF: 32,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 34,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []crust.CrustType{crust.CrustNone, crust.CrustSun, crust.CrustSun},
		},
    }

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc := tc
			res := crust.FindSunCrust(tc.data)
			assert.Equal(t, tc.expected, res)
		})
	}
}