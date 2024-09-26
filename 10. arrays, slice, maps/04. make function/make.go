package main

import "fmt"

func main() {
	// slices, maps, channels can be created using make function also
	// but it is only used for memory management scenarios
	// where you know that using make will increase efficiency of code.
	// In example of array, make() can pre allocate memory so that re-allocation
	// and de-allocation and dumping/recreating of arrays and slices can be reduces
	// Same goes for maps. But it is for very specific scenarios.

	// create a slice of strings using make() but with a initial length of 4 and maximum capacity
	// of 10 so that memory can be pre-allocated
	slice := make([]string, 4, 10)
	fmt.Println(len(slice), cap(slice)) // output 4, 10
	fmt.Println(slice)

	// now first 0-3 indexes of slice above are empty and can be allocated some values
	slice[0] = "1"
	slice[1] = "2"
	slice[2] = "3"
	slice[3] = "4"
	fmt.Println(slice)

	// now capacity of this slice is 10 (0-9 indexes) fill them also
	slice = append(slice, "5", "6", "7", "8", "9", "10")
	fmt.Println(slice)

	// create a map with make() with a capacity of 3
	mapMake := make(map[string]int, 3)
	mapMake["Zero"] = 0
	mapMake["One"] = 1
	mapMake["Two"] = 2
	fmt.Println(mapMake)

	// i can also break this capacity by easily adding one more element
	// but now this will cause re-creation of new map with re-allocation of
	// bigger memory, then placing all elements of previous map into new map
	// and so on.. because make() said capacity of 3 but you broke it by adding 4..
	mapMake["Break Capacity"] = 999
	fmt.Println(mapMake)
}

// So in front of screen nothing is special everything is just simple array/map operations
// but behind the scenes golang has pre-allocated some memory to this slice/map so that
// when items are deleted, added, more slices are created, etc. from this slice no
// re-allocation, dumping of memory is taking place.

// Because when you create a slice/map or delete a element or try to add more elements etc.
// if the capacity of the existing array/map is not enough.. then golang has to destroy
// the previous array/map.. create a new one with the big capacity.. and then add all these
// elements again and so on more and more things. This can be all diminished or highly reduced
// if make() is used and the scenario you are using make() in.. is right.

// This on a bigger level project can increase a lot of performance and give better
// memory management.

// Because make() tells go that this array, slice, channel will have this much capacity
// at most, even though you easily break this capacity by adding more elements normally
// but this pre allocation is very good idea if you know in advance that this array will
// have this much capacity and no more is needed.
// So in short when you know the capacity needed for a channel, array, map just use make()
// if it increases performance for your scenario.
