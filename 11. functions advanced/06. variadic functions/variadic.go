package main

import "fmt"

func main() {

	fmt.Println(addAll_V1(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(addAll_V2("i am k", 1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// 1. variation "one" of variadic function

// this variadic function accept unlimited "int"
func addAll_V1(nums ...int) int {

	// store the sum of all arguments here
	sum := 0

	// use a for loop to iterate over slice of arguments
	for _, value := range nums {
		sum = sum + value
	}

	return sum
}

// 2. variation "two" of variadic function

// this variadic function accept unlimited "int" but also accept one standalone argument
// called "k" which can be of any data type
// short answer it you can add as many standalone parameters you want
func addAll_V2(k string, nums ...int) int {

	// print the argument 1 for satisfaction
	fmt.Println(k)

	// store the sum of all arguments expect "k" here
	sum := 0

	// use a for loop to iterate over slice of arguments
	for _, value := range nums {
		sum = sum + value
	}

	return sum
}

// a variadic functions accepts infinite amount of parameters that are of same data type

// it can also accept standalone parameters of other data types also

// ...dots in front of a data type name are places in parameter so that this function
// become variadic and accept unlimited amount of parameters

// those infinite arguments (which are finite, infinite basically means any amount can be passed)
// which are passed to this function are put into the slice, so inside this variadic functions
// a slice needs to be handles, the slice that contains all the arguments that were passed to this
// function
