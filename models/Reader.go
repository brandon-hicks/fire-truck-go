package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Findtext() string {
	var text string
	fmt.Println("What calculation do you want to perform? Your choices are: gpm, pdp, FL, totalPL.")
	_, err := fmt.Scan(&text)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(text) == "" {
		log.Panic("Please enter a correct calculation type")
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
		log.Panic("Please enter a proper hose size")
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
		log.Panic("Please enter a proper length.")
	}
	hoseLength, err := strconv.Atoi(totalHoseLength)
	if err != nil {
		log.Fatal(err)
	}
	return hoseLength
}

func FindTipSize() float64 {
	var nozzleTipSize string
	fmt.Println("Please enter the tip size")
	_, err := fmt.Scan(&nozzleTipSize)
	if err != nil {
		log.Fatal(err)
	}
	if nozzleTipSize == "" {
		log.Panic("Please enter a proper tip size.")
	}
	tipSize, err := strconv.ParseFloat(nozzleTipSize, 64)
	if err != nil {
		log.Fatal(err)
	}
	return tipSize
}

func FindNozzleCoefficient() int {
	var nozzleType string
	fmt.Println("Please enter the nozzle coefficient. Your choices are: 50, 80, or 100")
	_, err := fmt.Scan(&nozzleType)
	if err != nil {
		log.Fatal(err)
	}

	switch nozzleType {
	case "50":
		return 50
	case "80":
		return 80
	case "100":
		return 100
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
