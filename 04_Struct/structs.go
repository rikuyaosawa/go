package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	fmt.Println("Hi, my name is", appUser.firstName)
	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
