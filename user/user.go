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

type Admin struct {
	email    string
	password string
	User     User
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
func NewAdmin(email string, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName:   "ADMIN",
			lastName:    "ADMIN",
			birthDate:   "---",
			createdTime: time.Now(),
		},
	}
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func (u *User) PrintUserDetails() {
	fmt.Print(u.firstName, u.lastName, u.birthDate, u.createdTime)
}
