package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFromConsole(textForConsole string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(textForConsole)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
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
	if nozzleType == "" {
		log.Panic("Please enter a proper coefficient.")
	}
	nozzleCoefficient, err := strconv.Atoi(nozzleType)
	if err != nil {
		log.Fatal(err)
	}
	return nozzleCoefficient
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
