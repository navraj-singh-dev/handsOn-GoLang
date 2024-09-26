package main

import (
	"fmt"
	"math"
)

func main() {

	// variables
	// one, two, three := 50, 50, 50
	// const one, two, three float64 = 50, 50, 50
	var numOne float64 = 50
	const numTwo float64 = 50
	numThree := 50.0

	result := (numOne + numTwo) * math.Pow(numThree, float64(2))

	// Taking input using scan
	var numFour float64
	fmt.Print("Give numFour: ")
	fmt.Scan(&numFour)
	result += numFour

	// Output to terminal
	// fmt.Print("\n", "The result is: ", result)
	// fmt.Println("The result is: ", result)
	// fmt.Printf("numOne value is: %v\nnumTwo value is: %v\nnumThree value is: %v\nnumFour value is: %v\nResult value is: %v\n ", numOne, numTwo, numThree, numFour, result)
	printResult := fmt.Sprint("\n", "The result is: ", result)
	fmt.Print(printResult)
}
