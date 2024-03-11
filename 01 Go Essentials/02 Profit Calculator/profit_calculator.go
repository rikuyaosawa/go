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

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println("EBT:", earningsBeforeTax)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}
