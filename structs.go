package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getInputFromUser("Enter the firstname:")
	userLastName := getInputFromUser("Enter the lastName:")
	userBirthDate := getInputFromUser("Enter the birthDate:")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	admin := user.NewAdmin("test@example.com", "test123")

	admin.User.PrintUserDetails()
	admin.User.ClearUserName()
	admin.User.PrintUserDetails()

	appUser.PrintUserDetails()
	appUser.ClearUserName()
	appUser.PrintUserDetails()
}

func getInputFromUser(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scanln(&input)
	return input
}
