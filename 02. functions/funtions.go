package main

import (
	"fmt"
)

func main() {
	numOne := getUserInput("Provide value of numOne: ")
	numTwo := getUserInput("Provide value of numTwo: ")
	numThree := getUserInput("Provide value of numThree: ")

	result := calculation(numOne, numTwo, numThree)

	fmt.Print("\n")
	fmt.Printf("The result is: %v", result)
	fmt.Print("\n")
}

func calculation(numOne, numTwo, numThree float64) (result float64) {
	result = numOne + numTwo + numThree
	return result
}

func getUserInput(text string) float64 {
	fmt.Print(text)

	// get the value
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}
