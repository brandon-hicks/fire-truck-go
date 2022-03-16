package models

import (
	"fmt"
	"strings"
)

func MakeCalculations() {
	text := Findtext()
	lowerCaseText := strings.ToLower(text)

	if lowerCaseText == "gpm" {
		fog := IsFogTip()
		tipSize := FindTipSize(fog)

		nozzleCoefficient := FindNozzleCoefficient(fog)

		fmt.Println(CalculateGallonsPerMinute(tipSize, float64(nozzleCoefficient), fog))

	} else if lowerCaseText == "pdp" {
		fog := IsFogTip()
		tipSize := FindTipSize(fog)
		nozzleCoefficient := FindNozzleCoefficient(fog)
		hoseLength := FindHoseLength()
		hoseSize := FindHoseSize()
		appliance := FindAppliance()

		fmt.Println(CalculatePumpDischargePressure(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance))

	} else if lowerCaseText == "fl" {
		fog := IsFogTip()
		tipSize := FindTipSize(fog)
		nozzleCoefficient := FindNozzleCoefficient(fog)
		hoseSize := FindHoseSize()
		hoseLength := FindHoseLength()

		fmt.Println(CalculateFrictionLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog))

	} else if lowerCaseText == "totalpl" {
		fog := IsFogTip()
		tipSize := FindTipSize(fog)
		nozzleCoefficient := FindNozzleCoefficient(fog)
		hoseLength := FindHoseLength()
		hoseSize := FindHoseSize()
		appliance := FindAppliance()

		fmt.Println(CalculateTotalPressureLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance))

	} else {
		fmt.Println("Please enter a valid calculation type.")
		MakeCalculations()
	}
}
