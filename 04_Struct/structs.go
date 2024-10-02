package main

import (
	"fmt"
	"time"
)

type user struct {
	// fields
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) sayHi() {
	fmt.Println("Hello! I am", u.firstName)
}

func (u *user) clearUser() {
	u.firstName = ""
	u.lastName = ""
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
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

	appUser.sayHi()
	appUser.clearUser()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
