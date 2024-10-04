package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	// fields
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

func New(fName, lName, bDate string) (User, error) {
	if fName == "" || lName == "" || bDate == "" {
		return User{}, errors.New("empty values are not accepted... exiting")
	}

	return User{
		firstName: fName,
		lastName:  lName,
		birthdate: bDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func (u User) SayHi() {
	fmt.Println("Hello! I am", u.firstName)
}

func (u *User) ClearUser() {
	u.firstName = ""
	u.lastName = ""
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}
