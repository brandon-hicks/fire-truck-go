package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindCalculationType() bool {
	var multiline string
	fmt.Println("Is this going to include multiple lines?")
	_, err := fmt.Scan(&multiline)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(multiline) == "yes" {
		return true
	} else if strings.ToLower(multiline) == "no" {
		return false
	}
	return FindCalculationType()
}

func Findtext() string {
	var text string
	fmt.Println("What calculation do you want to perform? Your choices are: gpm, pdp, FL, totalPL.")
	_, err := fmt.Scan(&text)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(text) == "" {
		fmt.Println("Please enter a correct calculation type")
		Findtext()
	}
	return text
}

func FindHoseSize() float64 {
	var totalHoseSize string
	fmt.Println("Please enter a hose size.")
	_, err := fmt.Scan(&totalHoseSize)
	if err != nil {
		log.Fatal(err)
	}
	if totalHoseSize == "" {
		fmt.Println("Please enter a proper hose size")
		FindHoseSize()
	}
	hoseSize, err := strconv.ParseFloat(totalHoseSize, 64)
	if err != nil {
		log.Fatal(err)
	}
	return hoseSize
}

func FindHoseLength() int {
	var totalHoseLength string
	fmt.Println("Please enter the hose length.")
	_, err := fmt.Scan(&totalHoseLength)
	if err != nil {
		log.Fatal(err)
	}
	if totalHoseLength == "" {
		fmt.Println("Please enter a proper length.")
		FindHoseLength()
	}
	hoseLength, err := strconv.Atoi(totalHoseLength)
	if err != nil {
		log.Fatal(err)
	}
	return hoseLength
}

func IsFogTip() bool {
	var fog string
	fmt.Printf("Yes/No: Is this a fog nozzle?")
	_, err := fmt.Scan(&fog)
	if err != nil {
		log.Fatal(err)
	}
	if strings.ToLower(fog) == "yes" {
		return true
	}
	return false
}

func FindTipSize(fog bool) float64 {
	var nozzleTipSize string
	if fog {
		return 1
	}
	fmt.Println("Please enter the tip size")
	_, err := fmt.Scan(&nozzleTipSize)
	if err != nil {
		log.Fatal(err)
	}
	if nozzleTipSize == "" {
		fmt.Println("You did not enter a proper tip size.")
		FindTipSize(fog)
	}
	tipSize, err := strconv.ParseFloat(nozzleTipSize, 64)
	if err != nil {
		fmt.Println("You did not enter a proper tip size.", FindTipSize(fog))
	}
	return tipSize
}

func FindNozzleCoefficient(fog bool) int {
	if fog {
		return 100
	}
	scanner := bufio.NewReader(os.Stdin)
	//var nozzleType string
	fmt.Println("Please enter nozzle type: Smooth Bore Hand or Smooth Bore Master")
	nozzleType, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	lowerCaseNozzleType := strings.ToLower(nozzleType)
	trimmedNozzleType := strings.Trim(lowerCaseNozzleType, "\n")

	switch trimmedNozzleType {
	case "smooth bore hand":
		return 50
	case "smooth bore master":
		return 80
	default:
		return 50
	}
}

func FindAppliance() bool {
	var appliance string
	fmt.Printf("Yes/No: Is there an appliance? ")
	_, err := fmt.Scan(&appliance)
	if err != nil {
		log.Fatal(err)
	}
	if appliance == "yes" {
		return true
	}
	return false
}

func FindElevation() bool {
	var elevation string
	fmt.Printf("Yes/No: Is there elevation? ")
	_, err := fmt.Scan(&elevation)
	if err != nil {
		log.Fatal(err)
	}
	if strings.ToLower(elevation) == "yes" {
		return true
	}
	return false
}

func FindElevationPsi(elevationGain string) int {
	if strings.ToLower(elevationGain) == "up" {
		return ElevationCalculation()
	} else if strings.ToLower(elevationGain) == "down" {
		return ElevationCalculation()
	}
	return 0
}

func ElevationCalculation() int {
	var typeOfElevation string
	var amountOfFloors float64
	var amountOfFeet float64
	fmt.Println("Are you adjusting for floors or hills?")
	_, err := fmt.Scan(&typeOfElevation)
	if err != nil {
		log.Fatal(err)
	}
	// for floors there is a 5 psi loss or gain per floor after the first floor.
	if strings.ToLower(typeOfElevation) == "floors" {
		fmt.Println("How many floors are you adjusting for?")
		_, err := fmt.Scan(&amountOfFloors)
		if err != nil {
			log.Fatal(err)
		}
		totalPressureLostFromFloors := (amountOfFloors - 1) * 5
		return int(totalPressureLostFromFloors)
	} else {
		fmt.Println("How many feet of elevation will you be adjusting for?")
		// for elevation that does not come from floors there is a .5 psi loss of pressure per foot over the pump and .5 psi gain for every foot under the pump.
		_, err := fmt.Scan(&amountOfFeet)
		if err != nil {
			log.Fatal(err)
		}
		totalPressureLostFromFeet := amountOfFeet * 0.5
		return int(totalPressureLostFromFeet)
	}
}

func FindElevationType() string {
	var elevationGain string

	fmt.Println("Are you going to be going up or down?")
	_, err := fmt.Scan(&elevationGain)
	if err != nil {
		log.Fatal(err)
	}
	return elevationGain
}
