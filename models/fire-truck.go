package models

import (
	"math"
)

func CoefficientDerivedFromHoseSize(hoseSize float64) float64 {
	switch hoseSize {
	case 0.75:
		return 1100.0
	case 1.5:
		return 24.0
	case 1.75:
		return 15.5
	case 2.5:
		return 2.0
	case 3.0:
		return 0.8
	case 4.0:
		return 0.2
	case 5.0:
		return 0.08
	default:
		return 15.5
	}
}

func CalculateGallonsPerMinute(tipSize float64, nozzle float64) float32 {
	return float32(29.7 * math.Pow(tipSize, 2) * math.Sqrt(nozzle))
}
