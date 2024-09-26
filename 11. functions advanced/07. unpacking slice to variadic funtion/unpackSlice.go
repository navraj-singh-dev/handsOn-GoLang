package main

import "fmt"

func main() {

	// a variadic function accepts as many arguments as one can give
	// but sometimes you might have a slice which contains all the arguments
	// you want to pass, but you dont want to type out all those elements
	// in the variadic function call.
	// so just pass the slice as it is in the function call definition
	// and then use 3 dots after the the name of the slice ...
	// this will unpack slice and put each element to the variadic function
	// as a standalone argument
	// example

	// you want this list's all elements to be passed as argument to the
	// variadic funtion
	slice := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 0}

	// just pass the slice and add ... after it, all its values will be unpacked
	// just think in terms of like the ... dots just passed all the values in the "slice"
	// variable as comma separated arguments automatically without you needing to type it out
	answer := addAll(slice...)
	fmt.Println(answer)
}

func addAll(s ...int) int {
	sum := 0

	for _, value := range s {
		sum += value
	}

	return sum
}
