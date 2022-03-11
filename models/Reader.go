package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func IsFogTip() bool {
	var fog string
	fmt.Printf("Yes/No: Is this a fog nozzle?")
	_, err := fmt.Scan(&fog)
	if err != nil {
		log.Fatal(err)
	}
	lowerCaseFog := strings.ToLower(fog)

	if lowerCaseFog == "yes" {
		return true
	}
	return false
}

func FindTipSize(fog bool) float64 {
	var nozzleTipSize string
	if fog != false {
		return 1
	}
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

func FindNozzleCoefficient(fog bool) int {
	if fog == true {
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
