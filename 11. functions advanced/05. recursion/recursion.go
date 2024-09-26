package main

import (
	"fmt"
)

func main() {

	// 1. recursive funtions call themselves in their own function body
	// 2. recursive functions have a exit out condition to prevent a
	//    infinite loop and out of memory

	// write a recursive funtion to find the factorial of a number
	// 5 * 4 * 3 * 2 * 1 -> factorial of 5, same for other
	answer := doRecursion(5)
	fmt.Println(answer)
}

func doRecursion(num int) int {
	if num == 0 {
		return 1
	}

	return num * doRecursion(num-1)
}

// this example is actually less performant when done with recursion
// a for loop will work better here.
// but its also a very easy example to understand recursion
// recursion is memory consuming and not the best option everytime.		