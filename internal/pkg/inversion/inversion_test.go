package inversion_test

import (
	"testing"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/pkg/inversion"
	"github.com/stretchr/testify/assert"
)

func TestFindInversionsAbove(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data []entity.WeatherStationData
		expected []bool
	}{
        {
			name: "no inversions",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 32,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 28,
				},
			},
			expected: []bool{false, false, false},
		},
		{
			name: "inversion between all elevations",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 28,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 100,
				},
			},
			expected: []bool{true, true, false},
		},
		{
			name: "inversion between lowest and middle elevations",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 28,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 25,
				},
			},
			expected: []bool{true, false, false},
		},
    }

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := inversion.FindInversionsAbove(tc.data)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestFindInversionsBelow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data []entity.WeatherStationData
		expected []bool
	}{
        {
			name: "no inversions",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 32,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 28,
				},
			},
			expected: []bool{false, false, false},
		},
		{
			name: "inversion between all elevations",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 28,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 100,
				},
			},
			expected: []bool{false, true, true},
		},
		{
			name: "inversion between lowest and middle elevations",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					TemperatureF: 28,
				},
				{
					ElevationFt: 500,
					TemperatureF: 30,
				},
				{
					ElevationFt: 1000,
					TemperatureF: 25,
				},
			},
			expected: []bool{false, true, false},
		},
    }

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := inversion.FindInversionsBelow(tc.data)
			assert.Equal(t, tc.expected, res)
		})
	}
}