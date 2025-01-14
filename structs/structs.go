package main

import (
	"errors"
	"fmt"
	"time"
	"example.com/structs/user"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	// ... do something awesome with that gathered data!
	if err != nil {
		fmt.Println(err)
		return
	}
	//outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
