package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	birthDate   string
	createdTime time.Time
}

func newUser(firstName string, lastName string, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Firstname, Last name and birth date are required")
	}

	return &user{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		createdTime: time.Now(),
	}, nil

}

func main() {
	userFirstName := getInputFromUser("Enter the firstname:")
	userLastName := getInputFromUser("Enter the lastName:")
	userBirthDate := getInputFromUser("Enter the birthDate:")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	if err == nil {
		fmt.Println(err)
		return
	}
	appUser.printUserDetails()
	appUser.clearUserName()
	appUser.printUserDetails()
}
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func getInputFromUser(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func (u *user) printUserDetails() {
	fmt.Printf(u.firstName, u.lastName, u.birthDate, u.createdTime)
}
