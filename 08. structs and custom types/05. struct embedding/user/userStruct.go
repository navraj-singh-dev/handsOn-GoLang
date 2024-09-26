// struct A can have a field which is a another struct B
// This is similar to inheritance in other langs.

// As struct A has struct B embedded in it all the methods
// and fields of the struct B are accessible by the struct A.
// Struct A can have its own methods and fields but including the
// struct B.

package user

import (
	"errors"
	"fmt"
	"time"
)

// This adming struct have its own fields as well as everything that user struct
// have. All methods and fields of user struct are in this Admin struct also
type Admin struct {
	email    string
	password string

	// Don't embed struct as "User User" bcuz then you have to do something like
	// user.User.<method name> which is long. Omit the "User" key name and directly
	// put the struct name then you can easily access methods like
	// user.<method name> without needing
	User
}

// constructor for Admin struct
func NewAdmin() Admin {

	var admin Admin = Admin{
		email:    "E-Mail",
		password: "Password",
		// As only "User" struct name is put as field in the Admin struct
		// during making of admin instance you must give the "User"
		// as key only, not any other name. See below i wrote "User"
		// as key name here not any other name
		User: User{
			userFirstName: "fname",
			userLastName:  "lname",
			userBirthData: "birth",
			createdAt:     time.Now(),
		},
	}

	return admin
}

type User struct {
	userFirstName string
	userLastName  string
	userBirthData string
	createdAt     time.Time
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields (firstName, lastName, birthdata) values are required")
	}

	var u *User = &User{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthData: birthdate,
		createdAt:     time.Now(),
	}

	return u, nil
}

func (u *User) PrintUserStruct() {
	fmt.Println((*u).userFirstName, u.userLastName, u.userBirthData)
	fmt.Println("This struct instance was created at: ", u.createdAt)
}

func (u *User) ClearFullName() {
	u.userFirstName = ""
	u.userLastName = ""
}
