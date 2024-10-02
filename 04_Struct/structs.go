package main

import (
	"errors"
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

func newUser(fName, lName, bDate string) (user, error) {
	if fName == "" || lName == "" || bDate == "" {
		return user{}, errors.New("empty values are not accepted... exiting")
	}

	return user{
		firstName: fName,
		lastName:  lName,
		birthdate: bDate,
		createdAt: time.Now(),
	}, nil
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

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.sayHi()
	appUser.clearUser()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
