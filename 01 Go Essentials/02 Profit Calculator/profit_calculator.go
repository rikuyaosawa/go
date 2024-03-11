package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("Ratio: %.3f", ratio)
}

func calculate(
	revenue,
	expenses,
	taxRate float64,
) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
