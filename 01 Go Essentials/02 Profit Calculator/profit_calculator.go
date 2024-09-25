package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile string = "result.txt"

func main() {
	fmt.Print("Revenue: ")
	revenue, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Print("Expenses: ")
	expenses, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Print("Tax Rate: ")
	taxRate, err := getInput()
	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("Ratio: %.3f", ratio)
	writeProfitToFile(profit)
}

func getInput() (float64, error) {
	var inputText float64
	fmt.Scan(&inputText)
	if inputText <= 0 {
		return 0, errors.New("0 or negative values are unacceptable")
	}
	return inputText, nil
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func writeProfitToFile(profit float64) {
	profitText := fmt.Sprint(profit)
	os.WriteFile(profitFile, []byte(profitText), 0o644)
}
