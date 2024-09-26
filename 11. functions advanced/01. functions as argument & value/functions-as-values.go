package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6}

	arr := transformNumber(sampleSlice, cubeNum)
	fmt.Println("transformed array cubed-num: ", arr)

	arr = transformNumber(sampleSlice, squareNum)
	fmt.Println("transformed array square-num: ", arr)
}

// function in go are first class values
// function can be set as a parameter of another function
// function can be given as a argument to another function

// create a function that takes another function as argument
// you must completely define the function type in the parameter excluding its name
// like what this parameter function will accept as its own arguments and what does it return itself
func transformNumber(array []int, function func(int) int) []int {
	// create a array to be returned by this function
	arr := []int{}

	// loop over array and get the values of its indexes
	for _, value := range array {
		// call the function which this function received as argument
		// and pass it the value which for loop is iterating on..
		tranformedValue := function(value)
		// append this value to the "arr" declared in this funtion
		arr = append(arr, tranformedValue)
	}

	return arr
}

// Both these functions below satisfies the transformNumber()'s 2nd parameter which is a function..
// requirements. Thus either of these two functions can be passed to transformNumber() at its 2nd parameter
func squareNum(num int) int {
	return num * num
}

func cubeNum(num int) int {
	return num * num * num
}

// Note: When defining the function to be set in the parameter, its definition can get very very long
// 			 for your scenarios. So consider using custom types (aliases).
