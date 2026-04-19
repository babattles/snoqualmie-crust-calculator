package entity_test

import (
	"testing"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestSortByElevation(t *testing.T) {
	t.Parallel()

	t.Run("Proof that sort actually sorts :P", func(t *testing.T) {
		t.Parallel()
		stationData := []entity.WeatherStationData{
			{ElevationFt: 50}, {ElevationFt: 10}, {ElevationFt: 30},
		}
		assert.Equal(t, 50, stationData[0].ElevationFt)

		entity.SortByElevation(stationData)
		assert.Equal(t, 10, stationData[0].ElevationFt)
	})
}
