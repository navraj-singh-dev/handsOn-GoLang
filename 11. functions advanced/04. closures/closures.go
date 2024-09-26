package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6}

	// multiplyWithTwo is that inner function being returned
	multiplyWithTwo := closureFunc(2) // 2 here will be "multiply parameter"

	// multiplyWithThree is that inner function being returned
	multiplyWithThree := closureFunc(3) // 3 here will be "multiply parameter"

	// multiply all array elements with 3 by passing the returned function
	arr_Multiplied_With_Two := transformNumber(sampleSlice, multiplyWithTwo)
	fmt.Println(arr_Multiplied_With_Two)

	// multiply all array elements with 3 by passing the returned function
	arr_Multiplied_With_Three := transformNumber(sampleSlice, multiplyWithThree)
	fmt.Println(arr_Multiplied_With_Three)


	// The cool thing is that the returned inner function still holds the same value of multiply
	// in it self, even though closureFunc was called and then it finished its job.
	// Example:
	// Here in -> multiplyWithTwo := closureFunc(2)
	// 2 here is the "multiply" parameter's value that is then used in returned inner function.
	// Now closureFunc(2) will run and then die, before dying it return me the anonymous function
	// All i am trying to say is that this retuned anonymous func still holds this "2" this multiply
	// parameter's value. It will always hold this value of 2.. it do not mean how many times 
	// you call this "multiplyWithTwo" (the returned inner func).
}

func transformNumber(array []int, function func(int) int) []int {
	arr := []int{}
	for _, value := range array {
		tranformedValue := function(value)
		arr = append(arr, tranformedValue)
	}
	return arr
}

func closureFunc(multiply int) func(int) int {
	// return a anonymous function
	return func(number int) int {
		// the "multiply" is a parameter of closureFunc()
		// but is used in this inner function
		return number * multiply
	}
}

// -- closures --

// 1. closures return a function as return value

// 2. this function locks-in/hold data from their scope which they are defined in.
//		like any normal function defined/declared outside of the main() are global scoped.
// 		thus they can access any variable or data defined in the global scope.
// 		like in my current closure the anonymous function is declared inside of another func,
//		thus this anonymous function's scope is "function scoped". thus this anonymous function
//		now hold every data which will be defined/declared in this specific "function scope".
// 		for example "multiply int" is parameter of closureFunc() but the anonymous function will
//		forever hold/lock-in this "multiply" parameter's value with it, because multiply is defined
// 		in the function scope and the anonymous function is also defined in that same scope.

// 3. So in short a closure is a pattern where you define a outer function which has it own data
//    (variables, parameter etc.). Then there is a inner function which then take holds of the
// 		all the parameters and variables of outer functions which are used in this inner function.
//		Then this inner function is returned by the outer function as return value.
//		Then one can use this returned function (inner function).
//		It is guarenteed that the returned function (inner function) still holds the data of outer
//		function which was used in it.
//
//		I hope it make funcking sense. It does to me right now!!
