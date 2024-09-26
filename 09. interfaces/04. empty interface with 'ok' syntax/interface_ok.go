package main

import (
	"fmt"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	intVal := 1
	floatVal := 1.1
	stringVal := "sting"

	printLimitedDataType(intVal)
	printLimitedDataType(floatVal)
	printLimitedDataType(stringVal)

	// create a custom value
	note := Note{
		Title:     "hello",
		Content:   "Content",
		CreatedAt: time.Now(),
	}
	// Now this function will take this as argument with no issue
	// but "ok" will not do anything about it and print
	// the default case
	printLimitedDataType(note)
}

// this function still accepts empty interface meaning it will accept any data-type
// as argument. but "ok" syntax will be used to only accept certain data-types only.
func printLimitedDataType(i interface{}) {

	// checking if argument is int
	intVal, ok := i.(int)
	// if it is then this if will run
	// i can also check for !ok as well if
	// that is more important
	if ok {
		fmt.Println("int: ", intVal)
		return
	}

	// check of argument is float
	floatVal, ok := i.(float64)
	if ok {
		fmt.Println("float: ", floatVal)
		return
	}

	// check if argument is string
	stringVal, ok := i.(string)
	if ok {
		fmt.Println("string: ", stringVal)
		return
	}

	fmt.Println("we dont support this unknown data-type in this function")
}
