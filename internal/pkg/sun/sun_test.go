package sun_test

import (
	"testing"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/pkg/sun"
	"github.com/stretchr/testify/assert"
)

func TestFindSunEffect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data []entity.WeatherStationData
		expected []bool
	}{
        {
			name: "all sun",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 500,
					RelativeHumidityPercent: 0,
				},
				{
					ElevationFt: 1000,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []bool{true, true, true},
		},
		{
			name: "no sun",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 500,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 1000,
					RelativeHumidityPercent: 100,
				},
			},
			expected: []bool{false, false, false},
		},
		{
			name: "uppermost sun",
			data: []entity.WeatherStationData{
				{
					ElevationFt: 0,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 500,
					RelativeHumidityPercent: 100,
				},
				{
					ElevationFt: 1000,
					RelativeHumidityPercent: 0,
				},
			},
			expected: []bool{false, false, true},
		},
    }

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := sun.FindSunEffect(tc.data)
			assert.Equal(t, tc.expected, res)
		})
	}
}