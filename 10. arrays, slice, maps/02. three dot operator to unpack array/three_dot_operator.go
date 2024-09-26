package main

import (
	"fmt"
)

func main() {

	// create a array
	arr := []string{}

	elementsToAdd := []string{"1", "2", "3", "4"}

	// append a "unpacked array" into array above
	arr = append(arr, elementsToAdd...) // ... dots after a array name unpacks it into multiple different values

	// print
	fmt.Println(arr)
}
