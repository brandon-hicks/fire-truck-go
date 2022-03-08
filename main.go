package main

import (
	"fire-truck-go/models"
	"fmt"
)

func main() {
	fmt.Println(models.CalculateActualGallonsPerMinute(1.25, 100))
	fmt.Println(models.CalculateFrictionLoss(300, 1.75, 0.88, 50))
	fmt.Println(models.CalculateTotalPressureLoss(300, 1.75, 0.88, 50, true))
	fmt.Println(models.CalculatePumpDischargePressure(300, 1.75, 0.88, 50, true))
}
