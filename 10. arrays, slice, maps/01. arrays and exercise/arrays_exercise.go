package main

import (
	"fmt"
)

func main() {

	// create array with 3 elements and print it
	arr3elements := [3]int{1, 2, 3}
	fmt.Println("array: ", arr3elements)

	// output only first element of array above
	first_element := arr3elements[0]
	fmt.Println("first element: ", first_element)

	// create a slice of index 2 and 3 of array above
	arrSlice := arr3elements[1:3]
	fmt.Println("slice with index 2 and 3: ", arrSlice)

	// create a slice of index 1 and 2 in two different slicing syntaxes
	arrSlice = arr3elements[:2]
	fmt.Println("Slice Syntax 1: ", arrSlice)
	arrSlice = arr3elements[0:2]
	fmt.Println("Slice Syntax 2: ", arrSlice)

	// modify the slice "exactly above" so that it has 2nd and last element of original array
	arrSlice = arrSlice[1:3]
	fmt.Println("Capacity: ", cap(arrSlice))
	fmt.Println("Modified Slice: ", arrSlice)

	// create new dynamic array with some elements
	dynamicArray := []string{"goal 1", "goal 2", "goal 3"}
	fmt.Println("dynamic array: ", dynamicArray)

	// change the index 1 value of this dynamic array and also append one element to last of it
	dynamicArray[1] = "goal modified"
	dynamicArray = append(dynamicArray, "appended goal")
	fmt.Println("dynamic array modified and append: ", dynamicArray)

	// create a struct then create a dynamic array then add the struct instances to this array
	type sampleStruct struct {
		name string
		game string
	}

	dynamicStructArray := []sampleStruct{
		{
			name: "nova 1",
			game: "hulu 2",
		},
		{
			name: "nova 2",
			game: "hulu 2",
		},
	}
	fmt.Println("dynamic stuct array: ", dynamicStructArray)

	dynamicStructArray = append(dynamicStructArray, sampleStruct{name: "nova 3", game: "hulu 3"})
	fmt.Println("dynamic struct array with append: ", dynamicStructArray)
}
