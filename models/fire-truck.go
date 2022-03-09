package models

import (
	"fmt"
)

func MakeCalculations() {
	text := Findtext()

	if text == "gpm" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()

		fmt.Println(CalculateActualGallonsPerMinute(tipSize, float64(nozzleCoefficient)))

	} else if text == "pdp" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()
		hoseLength := FindHoseLength()
		hoseSize := FindHoseSize()
		appliance := FindAppliance()

		fmt.Println(CalculatePumpDischargePressure(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), appliance))

	} else if text == "fl" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()
		hoseSize := FindHoseSize()
		hoseLength := FindHoseLength()

		fmt.Println(CalculateFrictionLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient)))

	} else if text == "totalpl" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()
		hoseLength := FindHoseLength()
		hoseSize := FindHoseSize()
		appliance := FindAppliance()

		fmt.Println(CalculateTotalPressureLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), appliance))

	} else {
		fmt.Println("You did not enter a correct calculation type.")
	}
}
