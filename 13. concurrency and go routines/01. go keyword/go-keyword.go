package main

import (
	"fmt"
	"time"
)

// create a function that takes 3 seconds to execute (simulate asyncronous task)
// simulate it using the time package
func greetWithTime(sec time.Duration) {
	time.Sleep(sec * time.Second) // here some thread will wait for 3 seconds, code below needs the this thread to be executed
	fmt.Printf("\nGreetings! After %v seconds", sec) // after waiting 3 sec thread can come here and execute this line
}

func main() {

	// using go routine for a function call
	go greetWithTime(3) // a go routine is a light weight-thread which will execute this function separate from main thread 

	// Now as i did not use "channels" this function did not output anything in terminal
	// because the main thread never waits for go routines to end, unless a channel is used.
	// 
	// So in next folder i will use "channels" and make this code work
	// Using channels in next folder this function's output will come in terminal.
}

// -- go keyword

//	It is used in front of a function call to offload a function to a lightweight go thread
//	which will execute this function in background concurrently/parallely. This keeps the main
//	thread free to keep the program responsive.
//
//	A function that has go keyword cannot return anything using "return" keyword. They need to
//	use channels to return anything from that function. But still "return" keyword can be used
//	inside go-keyworded-functions to exit the program. For example you "cannot" do "return <some value>"
//	inside of these functions and then receive this value when this function is called. But you can
//	use naked return keyword alone itself if you need to.
//
//	Go-keyworded-functions always and always need a buffered or unbuffered channel. Othwersie main
//	thread will end the program without waiting for other go routines to end. Channels block the
//	main thread so that it waits patiently for the go routines to end.
//
