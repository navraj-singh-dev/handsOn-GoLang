// A pointer is the actual ram address of a variable.
// We can directly mutate the value of that address without copying the variable.
// Go copies the variables we pass to any function and that is why
// we never mutate their actual value but the copied ones.
// So if a scenarios requires the actual original variable to be mutated
// and not a copy of it then pointers (&, *) must be used.

package main

import (
	"fmt"
)

func main() {

	var age1 int = 40
	var agePointer *int = &age1

	editPointerValue(agePointer)
	fmt.Println("The Pointer Of age: ", agePointer)
	fmt.Println("The age variable after giving its pointer to editPointerValue(): ", age1)

	// scan function from fmt is a great example, because we create a variable
	// with no value initially (eg: var age int) so that scan can store the user input in it.
	// But scan takes the pointer of this age variable like &age. Because here we actually need
	// to mutate the real age variable not the copied one.
	var age2 int
	fmt.Print("\nEnter your age: ")
	fmt.Scan(&age2)

	fmt.Println("The Pointer Of age2: ", &age2)
	editPointerValue(&age2)
	fmt.Println("The value of age2 after giving its pointer to editPointerValue: ", age2)
}

// This function takes a variable pointer and
// then edits the value at the address of the
// pointer. Mutating pointer.
func editPointerValue(pointer *int) {
	// * is used to get the actual value of the pointer
	// & -> address and * -> value at that address
	*pointer = *pointer - 18
}
