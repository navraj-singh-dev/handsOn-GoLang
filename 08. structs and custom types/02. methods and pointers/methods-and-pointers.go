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

// when a simple function in its receiver is derecing to a struct
// type, then that function is called a methods. Just like OOP and
// json objects. That function soley works for that struct type.

// Remember to use * with receiver parameter also, because go also just
// passes the copy of the struct if * is not used in receiver parameter
// if * is used in parameter of receiver then go automatically passes
// the struct object as a &pointer to the method and we need to write
// something like (*<struct name>).<method name> we can omit the (*) thing
func (u *user) printUserStruct() {
	fmt.Println((*u).userFirstName, u.userLastName, u.userBirthData)
	fmt.Println("This struct instance was created at: ", u.createdAt)
}

// A methods can also change the internal field data of the struct
func (u *user) clearFullName() {
	u.userFirstName = ""
	u.userLastName = ""
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var user1 user = user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthData: birthdate,
		createdAt:     time.Now(),
	}

	// using . we can also access the methods of a struct
	user1.printUserStruct()
	user1.clearFullName()
	user1.printUserStruct()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
