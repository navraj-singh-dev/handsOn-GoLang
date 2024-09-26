package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	userFirstName string
	userLastName  string
	userBirthData string
	createdAt     time.Time
}

// constructor funtions are not built-in go concept or primitive
// it is just a developer implementation and style of creatinga a
// constructor function inspired from other programming languages.
// this is why the constructor funtion also returs a copy of the
// struct not the original one as it is same as simple functions.
// So use & and * wherever required in constructor func also.
func newUser(firstName, lastName, birthdate string) (*user, error) {

	// the two benefits of using constructor func is :
	// 1. create your struct object with more speed
	// 2. add validations before creating the struct

	// validation using conditionals, one of the usecases of constructor funtions
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields (firstName, lastName, birthdata) values are required")
	}

	var u *user = &user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthData: birthdate,
		createdAt:     time.Now(),
	}

	return u, nil
}

func (u *user) printUserStruct() {
	fmt.Println((*u).userFirstName, u.userLastName, u.userBirthData)
	fmt.Println("This struct instance was created at: ", u.createdAt)
}

func (u *user) clearFullName() {
	u.userFirstName = ""
	u.userLastName = ""
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1, err := newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	user1.printUserStruct()
	user1.clearFullName()
	user1.printUserStruct()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// reason to shift from Scan() -> Scanln() is that in console
	// even if i press enter i wont go to next prompt in case of Scan()
	// it will always keep me with the first prompt even if i press enter 1000 times.
	// Scanln() solves this problem and let us enter in console if i do not want to
	// enter a value for a Scanln() prompt.
	fmt.Scanln(&value)
	return value
}
