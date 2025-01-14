package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	//pointers to structs in go don't need dereferencing
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() { //when editing variables related to structs, need to use pointer
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) { //constructor function
	//constructor functions are useful if you want to run some validation logic when creating struct
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are requried")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
