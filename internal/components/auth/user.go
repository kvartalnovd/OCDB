package auth

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	DateBirth time.Time
}

func NewUser(firstName, secondName, email string, dateBirth time.Time) *User {
	return &User{
		FirstName: firstName,
		LastName:  secondName,
		Email:     email,
		DateBirth: dateBirth,
	}
}

func (u *User) Show() {
	fmt.Printf("%s - %s\n", u.FirstName, u.LastName)
}
