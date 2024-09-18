package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0o644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func main() {
	isContinue := true
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
	}

	for isContinue {
		fmt.Println("\nWelcome to Go Bank!")
		fmt.Println("What do you wat to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var userChoice int
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Print("Your balance: ", balance)
		} else if userChoice == 2 {
			var deposit float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&deposit)
			balance += deposit
			writeBalanceToFile(balance)
		} else if userChoice == 3 {
			var withdraw float64
			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdraw)
			if balance < withdraw {
				fmt.Println("You cannot withdraw more than you have in balance.")
			} else {
				balance -= withdraw
				writeBalanceToFile(balance)
			}
		} else if userChoice == 4 {
			fmt.Println("Goodbye!")
			isContinue = false
		} else {
			fmt.Print("Invalid input. Please try again.")
		}
	}
}
