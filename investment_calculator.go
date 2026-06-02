package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

// go run .
func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:", futureValue)

	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Printf('Future Value: %.2f
	// Future Value (adjusted for inflation): %.2f', futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}
