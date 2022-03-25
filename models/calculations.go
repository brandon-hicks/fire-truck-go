package models

import (
	"math"
	"strings"
)

// Finds the value of C in the Friction Loss Formula
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

// GPM=29.7*(TP^2)*(sqrt NP)
// Finds the value of Q in the Friction Loss Formula and devides it by 100
// If the user is using a fog nozzle then the calculation returns 100. This is because this is the GPM of fog nozzles.
func CalculateGallonsPerMinute(tipSize float64, nozzle float64, fog bool) int {
	if !fog {
		calculation := float64(29.7 * math.Pow(tipSize, 2) * math.Sqrt(nozzle))
		return int(calculation)
	}
	return 100
}

// FL =C*(Q/100)^2*L/100
func CalculateFrictionLoss(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool) int {
	calculation := float64(CoefficientDerivedFromHoseSize(hoseSize) * math.Pow(float64(CalculateGallonsPerMinute(tipSize, nozzle, fog)/100), 2) * hoseLength / 100)
	return int(calculation)
}

// If an appliance is added then this will add 10 to that calculation. There is still some work to be done around this part.
func CalculateForAppliance(appliance bool) int {
	if !appliance {
		return 0
	}
	return 10
}

func CalculateForElevation(elevation bool, elevationType string) int {
	if !elevation {
		return 0
	}
	elevationNumber := FindElevationPsi(elevationType)
	return elevationNumber

}

// TPL=FL+AL+/-EP
func CalculateTotalPressureLoss(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool, appliance bool, elevation bool) int {
	if elevation {
		elevationType := FindElevationType()
		if strings.ToLower(elevationType) == "down" {
			calculation := float64(CalculateFrictionLoss(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance)-CalculateForElevation(elevation, elevationType))
			return int(calculation)
		}
		calculation := float64(CalculateFrictionLoss(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance)+CalculateForElevation(elevation, elevationType))
		return int(calculation)
	}
	calculation := float64(CalculateFrictionLoss(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance))
	return int(calculation)
}

// PDP=NP+TPL
func CalculatePumpDischargePressure(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool, appliance bool, elevation bool) int {
	calculation := float64(CalculateTotalPressureLoss(hoseLength, hoseSize, tipSize, nozzle, fog, appliance, elevation)) + nozzle
	return int(calculation)
}

// FL =(C*(Q/100)^2*L/100)/2
func CalculateFrictionLossInMultipleLinesOfSameSize(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool) int {
	calculation := (CalculateFrictionLoss(hoseLength, hoseSize, tipSize, nozzle, fog)) / 2
	return int(calculation)
}

// TPL=FL+AL+/-EP
func CalculateTotalPressureLossInMultiLinesOfSameSize(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool, appliance bool, elevation bool) int {
	if elevation {
		elevationType := FindElevationType()
		if strings.ToLower(elevationType) == "down" {
			calculation := float64(CalculateFrictionLossInMultipleLinesOfSameSize(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance)-CalculateForElevation(elevation, elevationType))
			return int(calculation)
		}
		calculation := float64(CalculateFrictionLossInMultipleLinesOfSameSize(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance)+CalculateForElevation(elevation, elevationType))
		return int(calculation)
	}
	calculation := float64(CalculateFrictionLossInMultipleLinesOfSameSize(hoseLength, hoseSize, tipSize, nozzle, fog)) + float64(CalculateForAppliance(appliance))
	return int(calculation)
}

// PDP=NP+TPL
func CalculatePumpDischargePressureInMultiLinesOfSameSize(hoseLength float64, hoseSize float64, tipSize float64, nozzle float64, fog bool, appliance bool, elevation bool) int {
	calculation := float64(CalculateTotalPressureLossInMultiLinesOfSameSize(hoseLength, hoseSize, tipSize, nozzle, fog, appliance, elevation)) + nozzle
	return int(calculation)
}
