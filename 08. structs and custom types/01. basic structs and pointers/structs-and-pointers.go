package main

import (
	"fmt"
	"time"
)

type user struct {
	userFirstName string
	userLastName  string
	userBirthData string
	createdAt     time.Time
}

func main() {
	// these scanned input values will be stored in the struct
	// this allow setting the values to struct field dynamically
	// by allowing user to input the data in console 
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// structs to me are very similar to json object or OOP
	// you create a blueprint called struct and then its instances
	// can be created.

	// structs allow to add functions, variables etc. of different types in it.
	// you create fields and then these fields can be filled with any built-in
	// data types or custom data types like interfaces and structs you create
	// or can be filled with functions, very similar to json or OOP
	var user1 user = user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthData: birthdate,
		createdAt:     time.Now(),
	}

	printUserStruct(&user1)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

// in structs we dont need to use * to get its pointer value
// because go allows this as shortcut, we can simply give name
// of struct and use . to access the fields of the original struct
// without needing to use * symbol in front of the variable name that
// hold the struct.
// But still if you want to use * then you can use it by (*<name>) syntax
func printUserStruct(u *user) {
	fmt.Println((*u).userFirstName, u.userLastName, u.userBirthData)
	fmt.Println("This struct instance was created at: ", u.createdAt)
}
