// A single struct with its methods and data can be put into a single standalone
// package of its own. Now this package can be used in other packages across the moduele.

// But creating a new package for structs and its methods|data can create new challenges
// remember the captial letter to export anything out of the package (struct, methods, variables etc.)

// structs when exported outside of package do not export their fields 😐
// to expose their fields one must export their fields also.. use the capital letter
// for every field you want to export. It goes the same for methods and variables and any other things.

// On a side note do not export the struct's fields just use a constructor funtion to return the 
// struct instance instead.

// In packages the constructor funtion is just named "New". Because in other packages using the
// consturctor function as user.NewUser() look just weird but user.New() looks good. So every golang
// develoer just names their constructor as "New" for every package they might create.

// In other packages the constructor function is used to create the struct instance, the
// old way of refering to the name of the struct and then adding value do not work as fields
// are probhibited to exporting.

package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	// to export these fields below
	// capitalize their first character
	// otherwise they cannot be exported
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
