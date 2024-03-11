package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// Alternative way: investmentAmout, expectedReturnRate, years := 1000.0, 5.5, 10.0
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(
		investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %v\n", futureValue)
	fmt.Printf("Future Real Value: %v\n", futureRealValue)
}

func calculateFutureValues(
	investmentAmount,
	expectedReturnRate,
	years float64) (float64, float64) {
	fv := (investmentAmount * math.Pow(1+expectedReturnRate/100, years))
	frv := (fv / math.Pow(1+inflationRate/100, years))
	return fv, frv
}

// Alternative Way //
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
// 	fv = (investmentAmount * math.Pow(1+expectedReturnRate/100, years))
// 	frv = (fv / math.Pow(1+inflationRate/100, years))
// 	return
// }
