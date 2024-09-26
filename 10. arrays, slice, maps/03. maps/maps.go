package main

import "fmt"

func main() {

	// create a map with keys of type string and values also of type string
	mapString := map[string]string{
		"keyOne": "Value One",
		"keyTwo": "Value Two",
	}
	fmt.Println(mapString)

	// keys data-type and value's data type can be anything (struct, int, float bool)
	mapInt := map[int]float64{
		1: 11.1,
		2: 12.2,
	}
	fmt.Println(mapInt)

	// adding a value to the map
	// if the key do not exist then it will be appended to the map
	mapInt[3] = 13.3
	fmt.Println(mapInt)

	// modfying a existing key's value
	mapInt[3] = 999.999
	fmt.Println(mapInt)

	// delete a key
	delete(mapInt, 3)
	fmt.Println(mapInt)
}
