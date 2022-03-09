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

func CalculateActualGallonsPerMinute(tipSize float64, nozzle float64) int {
	calculation := float64(29.7 * math.Pow(tipSize, 2) * math.Sqrt(nozzle))
	return int(calculation)
}

func CalculateGallonsPerMinute(tipSize float64, nozzle float64) int {
	calculation := float64(29.7 * math.Pow(tipSize, 2) * math.Sqrt(nozzle) / 100)
	return int(calculation)
}

func CalculateFrictionLoss(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64) int {
	calculation := float64(CoefficientDerivedFromHoseSize(hoseSize) * math.Pow(float64(CalculateGallonsPerMinute(tipSize, nozzle)), 2) * (hoseLength / 100))
	return int(calculation)
}

func CalculateForAppliance(appliance bool) int {
	if !appliance {
		return 0
	}
	return 10
}

func CalculateTotalPressureLoss(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, appliance bool) int {
	calculation := float64(CalculateFrictionLoss(hoseLength, hoseSize, tipSize, nozzle)) + float64(CalculateForAppliance(appliance))
	return int(calculation)
}

func CalculatePumpDischargePressure(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, appliance bool) int {
	calculation := float64(CalculateTotalPressureLoss(hoseLength, hoseSize, tipSize, nozzle, appliance)) + nozzle
	return int(calculation)
}
