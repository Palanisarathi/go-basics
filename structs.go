package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	birthDate   string
	createdTime time.Time
}

func newUser(firstName string, lastName string, birthDate string) *user {
	return &user{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		createdTime: time.Now(),
	}

}

func main() {
	userFirstName := getInputFromUser("Enter the firstname:")
	userLastName := getInputFromUser("Enter the lastName:")
	userBirthDate := getInputFromUser("Enter the birthDate:")

	var appUser *user
	appUser = newUser(userFirstName, userLastName, userBirthDate)
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
	fmt.Scan(&input)
	return input
}

func (u *user) printUserDetails() {
	fmt.Printf(u.firstName, u.lastName, u.birthDate, u.createdTime)
}
