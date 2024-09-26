package main

import (
	"fmt"

	"e.com/package-struct-export/user"
)

func main() {

	// I create the admin instance, i cannot get email, password by using '.'
	// because they are not exported but User struct which is embedded in the
	// admin struct is exported so i can access its methods completely.
	fmt.Println()
	fmt.Println("-- Admin Struct --")
	adminUser := user.NewAdmin()
	adminUser.PrintUserStruct()
	fmt.Println(adminUser.User)
}
