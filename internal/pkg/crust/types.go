package crust

import "github.com/babattles/snoqualmie-crust-calculator/internal/entity"

type CrustType string

const (
	CrustSun CrustType = "sun"
	CrustSunMaybe CrustType = "sunMaybe"
	CrustRain CrustType = "rain"
	CrustMelt CrustType = "melt"
	CrustNone CrustType = "none"
)

var crustPriorities []CrustType = []CrustType{
	CrustSun,
	CrustSunMaybe,
	CrustRain,
	CrustMelt,
	CrustNone,
}

func (c CrustType) GetPriority() int {
	for i, elm := range crustPriorities {
		if c == elm {
			return i
		}
	}
	return entity.MaxInt
}

func (c CrustType) Trumps(c2 CrustType) bool {
	return c.GetPriority() < c2.GetPriority()
}