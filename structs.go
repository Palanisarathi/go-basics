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

func main() {
	userFirstName := getInputFromUser("Enter the firstname:")
	userLastName := getInputFromUser("Enter the lastName:")
	userBirthDate := getInputFromUser("Enter the birthDate:")

	var appUser user
	appUser = user{
		firstName:   userFirstName,
		lastName:    userLastName,
		birthDate:   userBirthDate,
		createdTime: time.Now(),
	}
	printUserDetails(appUser)
}

func getInputFromUser(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scan(&input)
	return input
}

func printUserDetails(u user) {
	fmt.Printf("firstName: %v LastName:%v birthDate:%v createdTime:%v", u.firstName, u.lastName, u.birthDate, u.createdTime)
}
