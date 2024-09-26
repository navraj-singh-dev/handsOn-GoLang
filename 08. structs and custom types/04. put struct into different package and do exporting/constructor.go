package main

import (
	"fmt"
	// import the user package use its things by using syntax "user.<resource name>"
	"e.com/package-struct-export/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// refer to anything in other package by using it name the "." notation
	user1, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// this syntax will not work to create a user struct instance
	// because its struct's keys are not exported also its not
	// recommended to export the keys of the struct from other
	// packages
	// user2 = user.User{
	// 	userFirstName: "Hello",
	// }

	user1.PrintUserStruct()
	user1.ClearFullName()
	user1.PrintUserStruct()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
