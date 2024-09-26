package main

import (
	"bufio"
	"fmt"
	"os"
)

// open a file and close it with defer
func openAndClose() {

	file, err := os.Open("open_me_close_me.txt")

	if err != nil {
		fmt.Println(err)
	}
	// this defer statement will only execute when openAndClose() will return/ends
	defer file.Close() // yes this statement executes at absolute-last even though it is defined above some code

	// read the file by making scanner object
	scanner := bufio.NewScanner(file)
	// scan the first line of .txt file
	scanner.Scan()
	// return the scanned line as text
	fileContent := scanner.Text()
	// print the text
	fmt.Println(fileContent)

	// now here openAndClose() code ends and it returns, just when it return defer executes
	// file.Close() and BOOM file is closed.
}

func main() {

	openAndClose()
}

// -- defer keyword

//	defer keyword is used to execute a function. But it executes that function only when the function
//	enclosing defer keyword returns.
//
//	Meaning if i am using defer in a function named HELLO() then HELLO() is the function enclosing
//	defer keyword. So when HELLO() will return at that time defer will execute the function it is
//	supposed to run.
//
//	If defer is openly defined in the main() then main() is the enclosing function and when main()
//	returns, defer executes the function it is supposed to run.
//
//	If multiple 'defer' statements are declared in any function then the last/ending defer statement is
//	executed first, or to say when multiple defer statements are there LIFO acts, Last In First Out.
//	So defer statements in 1,2,3,4 order will be excuted as 4,3,2,1. I hope it makes sense
//
//	defer is mostly used to close file or close channels or any kind of suitable cleanup tasks, but still
//	it depends on you how you can use it for your scenario.
//
//
