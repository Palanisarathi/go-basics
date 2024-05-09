package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthDate   string
	createdTime time.Time
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Firstname, Last name and birth date are required")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		createdTime: time.Now(),
	}, nil

}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func (u *User) PrintUserDetails() {
	fmt.Print(u.firstName, u.lastName, u.birthDate, u.createdTime)
}
