package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/pkg/crust"
	"github.com/babattles/snoqualmie-crust-calculator/internal/repo/webapi/snowobs"
)

var ErrUnknownMountain = errors.New("unknown mountain")

type CrustLayer struct {
	ElevationFt int
	SunCrust    crust.CrustType
	MeltCrust   crust.CrustType
}

type MountainCrusts struct {
	Mountain entity.MountainType
	Layers   []CrustLayer
}

type Usecase struct {
	snowobs *snowobs.Client
}

func New(s *snowobs.Client) *Usecase {
	return &Usecase{snowobs: s}
}

func (u *Usecase) GetCrustsForMountain(ctx context.Context, mountain entity.MountainType) (MountainCrusts, error) {
	m, ok := entity.GetMountain(mountain)
	if !ok {
		return MountainCrusts{}, fmt.Errorf("%w: %s", ErrUnknownMountain, mountain)
	}

	end := time.Now()
	start := end.Add(-24 * time.Hour)

	resp, err := u.snowobs.GetStationData(ctx, m.StIDs, start, end)
	if err != nil {
		return MountainCrusts{}, fmt.Errorf("fetch station data: %w", err)
	}

	layers := make([]entity.WeatherStationData, 0, len(resp.Stations))
	for _, s := range resp.Stations {
		obs := s.Observations
		n := len(obs.DateTime)
		if n == 0 || len(obs.AirTemp) < n || len(obs.RelativeHumidity) < n {
			continue
		}
		i := n - 1
		layers = append(layers, entity.WeatherStationData{
			ElevationFt:             int(s.Elevation),
			TemperatureF:            int(obs.AirTemp[i]),
			RelativeHumidityPercent: int(obs.RelativeHumidity[i]),
		})
	}
	entity.SortByElevation(layers)

	sunCrusts := crust.FindSunCrust(layers)
	meltCrusts := crust.FindMeltCrust(layers)

	out := MountainCrusts{
		Mountain: mountain,
		Layers:   make([]CrustLayer, len(layers)),
	}
	for i, l := range layers {
		out.Layers[i] = CrustLayer{
			ElevationFt: l.ElevationFt,
			SunCrust:    sunCrusts[i],
			MeltCrust:   meltCrusts[i],
		}
	}
	return out, nil
}
