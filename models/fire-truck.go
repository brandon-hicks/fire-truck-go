package models

import (
	"fmt"
	"strings"
)

func MakeCalculations() {
	multiline := FindCalculationType()

	if !multiline {
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
			elevation := FindElevation()

			fmt.Println(CalculatePumpDischargePressure(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance, elevation))

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
			elevation := FindElevation()

			fmt.Println(CalculateTotalPressureLoss(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance, elevation))

		} else {
			fmt.Println("Please enter a valid calculation type.")
			MakeCalculations()
		}
	} else {
		text := Findtext()

		if strings.ToLower(text) == "fl" {
			fog := IsFogTip()
			tipSize := FindTipSize(fog)
			nozzleCoefficient := FindNozzleCoefficient(fog)
			hoseSize := FindHoseSize()
			hoseLength := FindHoseLength()

			fmt.Println(CalculateFrictionLossInMultipleLinesOfSameSize(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog))

		} else if strings.ToLower(text) == "totalpl" {
			fog := IsFogTip()
			tipSize := FindTipSize(fog)
			nozzleCoefficient := FindNozzleCoefficient(fog)
			hoseLength := FindHoseLength()
			hoseSize := FindHoseSize()
			appliance := FindAppliance()
			elevation := FindElevation()

			fmt.Println(CalculateTotalPressureLossInMultiLinesOfSameSize(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance, elevation))
		} else if strings.ToLower(text) == "pdp" {
			fog := IsFogTip()
			tipSize := FindTipSize(fog)
			nozzleCoefficient := FindNozzleCoefficient(fog)
			hoseLength := FindHoseLength()
			hoseSize := FindHoseSize()
			appliance := FindAppliance()
			elevation := FindElevation()

			fmt.Println(CalculatePumpDischargePressureInMultiLinesOfSameSize(float64(hoseLength), hoseSize, tipSize, float64(nozzleCoefficient), fog, appliance, elevation))
		} else if strings.ToLower(text) == "gpm" {
			fog := IsFogTip()
			tipSize := FindTipSize(fog)

			nozzleCoefficient := FindNozzleCoefficient(fog)

			fmt.Println(CalculateGallonsPerMinute(tipSize, float64(nozzleCoefficient), fog))

		} else {
			fmt.Println("Please enter a valid calculation type.")
			MakeCalculations()
		}
	}
}
