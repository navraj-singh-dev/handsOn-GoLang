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
	// but switch statemtn will not do anything about it and print
	// the default case
	printLimitedDataType(note)
}

// this function still accepts empty interface meaning it will accept any data-type
// as argument. but switch-type will be used to only accept certain data-types only.
func printLimitedDataType(i interface{}) {

	// now only
	switch i.(type) {
	case int:
		fmt.Println("Int Detected: ", i)
	case string:
		fmt.Println("String Detected: ", i)
	case float64:
		fmt.Println("Flaot64 Detected: ", i)
	default:
		fmt.Println("We do not support this unknown data type")
	}
}
