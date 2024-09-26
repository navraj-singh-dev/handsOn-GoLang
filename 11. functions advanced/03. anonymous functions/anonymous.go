package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6}

	// anonymous functions are nameless function which are "never" declared as a normal function and
	// then called. Instead they are nameless and directly implemented and passed or returned 
	// where-ever required. 
	// For example, a anonymous function can be passed as argument to another function directly in the
	// function's  call decalaration.
	// example can be:
	arr := transformNumber(sampleSlice, func(num int) int {
		return num * num
	})
	fmt.Println("squareFunc: ", arr)

	// see i dont need to go outside main() and then there declare a normal function 
	// and then pass it as argument into this func below. I can directly skip these steps
	// and pass a anonymous function here directly
	arr = transformNumber(sampleSlice, func(num int) int {
		return num * num * num
	})
	fmt.Println("cubeFunc: ", arr)

}

func transformNumber(array []int, function func(int) int) []int {
	arr := []int{}
	for _, value := range array {
		tranformedValue := function(value)
		arr = append(arr, tranformedValue)
	}
	return arr
}

// -- anonymous function --
// When to use? Well if a function is only gonna be passed as argument or returned as value
// once or twice and is not going to be used regularly by your code then you can simply skip
// declaring this function outside of main() and giving it a name and then passing it as argument
// to another function or returning it as a value from another funtion.

// You can simply go to the function in main() which wants a func as its argument and pass the
// anonymous function as argument there. No need to declare this function outside of main and
// giving it a name and then passing it as a argument in main() inside another function's call
// definition.

// Anonymous functions are good if your function is very short and only gonna be used one or twice.
// Other wise stick to normal named functions.
