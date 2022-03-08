package models

import (
	"fmt"
)

func MakeCalculations() {
	text := ReadFromConsole("What calculation do you want to perform? Your choices are: gpm, pdp, FL, totalPL. ")

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

	} else if text == "FL" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()
		hoseSize := FindHoseSize()
		hoseLength := FindHoseLength()

		fmt.Println(CalculateFrictionLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient)))

	} else if text == "totalPL" {
		tipSize := FindTipSize()
		nozzleCoefficient := FindNozzleCoefficient()
		hoseLength := FindHoseLength()
		hoseSize := FindHoseSize()
		appliance := FindAppliance()

		fmt.Println(CalculateTotalPressureLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), appliance))

	} else {
		fmt.Println("Please enter one of the following to get a proper calculation: gpm, pdp, FL, totalPL.")
	}
}
