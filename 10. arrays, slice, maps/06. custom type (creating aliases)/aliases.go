package main

import "fmt"

func main() {

	// sometimes defining a definition of map or array or function can get very long
	// like: func name([]string, map[string][]int, map[int]string) (int, string, float64)
	// like: map[string][]string
	// etc. Now to create a nickname for these definitions you created like the one above
	// you can use custom types meaning you can give them a nickname which in programming is
	// called aliases.
	// Meaning you wont need to type long definitions like:
	// func name([]string, map[string][]int, map[int]string) (int, string, float64)
	// you will just call their nickname and go will internally refer to the code definition.

	// create a alias for long function definition
	type spiderMan func([]string, map[string][]int, map[int]string) (int, string, float64)
	// then you will call this function simply by its nickname and pass the required arguments
	// like : spiderMan(<arguments here>) {..logic here}
	// i cannot implement this function here because it is merely a made-up example and it will take
	// enormous amount of time to just complete it out and completing it out is not any help
	// to the concept of "custom type (aliases)" which i am explaning.

	// create a alias for a map definition
	type strMap map[string]string
	// then use the alias name without needing to type " map[string]string "
	mapSimple := strMap{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}
	fmt.Println(mapSimple)

	// line break
	fmt.Println()
	fmt.Println()

	// make function can also work with aliases if they are aliases of array/map/channel
	mapMake := make(strMap, 2)
	mapMake["map"] = "make"
	mapMake["map2"] = "make2"
	fmt.Println(mapMake)

}
