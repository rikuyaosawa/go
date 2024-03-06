package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	// Alternative way: investmentAmout, expectedReturnRate, years := 1000.0, 5.5, 10.0
	var investmentAmout, years, expectedReturnRate float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmout)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmout * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}
