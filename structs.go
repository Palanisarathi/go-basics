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
	firstName := getInputFromUser("Enter the firstname:")
	lastName := getInputFromUser("Enter the lastName:")
	birthDate := getInputFromUser("Enter the birthDate:")
	createdTime := getInputFromUser("Enter the createdTime:")

	printUserDetails(firstName, lastName, birthDate, createdTime)
}

func getInputFromUser(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scan(&input)
	return input
}

func printUserDetails(firstName, lastName, birthDate, createdTime string) {
	fmt.Printf("firstName: %v LastName:%v birthDate:%v createdTime:%v", firstName, lastName, birthDate, createdTime)
}
