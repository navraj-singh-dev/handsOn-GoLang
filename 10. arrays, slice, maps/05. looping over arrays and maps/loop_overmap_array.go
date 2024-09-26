package main

import "fmt"

func main() {

	// loop over the array/slice
	arr := []string{"one", "two", "three", "four", "five"}

	// range is very important key word in loop
	// it set the for-loop to loop till all the elements of your array/map are
	// iterated and only then it stops
	// "range" returns two things: the index and the index's value if it is array
	// "range" returns two things: the key and the key's value if it is map
	for index, value := range arr {
		fmt.Println(index)
		fmt.Println(value)
	}

	// line breaks
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// loop over map with range
	mapSimple := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	for key, value := range mapSimple {
		fmt.Println(key)
		fmt.Println(value)
	}
}
