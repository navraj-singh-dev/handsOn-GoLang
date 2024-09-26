package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6}

	// return a function from a function and store them in variables
	squareFunc := returnFunc("square")
	cubeFunc := returnFunc("cube")

	// pass these two functions alternatively to the transformNumber()
	arr := transformNumber(sampleSlice, squareFunc)
	fmt.Println("squareFunc: ", arr)

	arr = transformNumber(sampleSlice, cubeFunc)
	fmt.Println("cubeFunc: ", arr)

}

// this function return another function as a return value
// this is a sample example of just how this returned function can be used
func returnFunc(func_name string) func(int) int {
	if func_name == "square" {
		return squareNum
	} else if func_name == "cube" {
		return cubeNum
	}
	return squareNum
}

func transformNumber(array []int, function func(int) int) []int {
	arr := []int{}
	for _, value := range array {
		tranformedValue := function(value)
		arr = append(arr, tranformedValue)
	}
	return arr
}

// One of these two function below will be returned according to argument passed
func squareNum(num int) int {
	return num * num
}

func cubeNum(num int) int {
	return num * num * num
}

// My example is bit of weird and just a extra step but its fine for me
// Just remember a function can return another function as return value,
// also accept another function as argument if its parameters allows so.