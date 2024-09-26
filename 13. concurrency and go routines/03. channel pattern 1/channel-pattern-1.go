package main

import (
	"fmt"
	"time"
)

// These 3 functions are 100% same just name is different
// All 3 of them will use the same channel
// All 3 of these will input value to the same channel
func greetWithTime_1(sec time.Duration, channel chan bool) {
	time.Sleep(sec * time.Second)
	fmt.Printf("\nGreetings _1 ! After %v seconds", sec)
	channel <- true
}
func greetWithTime_2(sec time.Duration, channel chan bool) {
	time.Sleep(sec * time.Second)
	fmt.Printf("\nGreetings _2 ! After %v seconds", sec)
	channel <- true
}
func greetWithTime_3(sec time.Duration, channel chan bool) {
	time.Sleep(sec * time.Second)
	fmt.Printf("\nGreetings _3 ! After %v seconds", sec)
	channel <- true
}

func main() {

	// -- Pattern 1 (not recommended at all, but still good to know)

	// It is just about making only one channel.. but making multiple go routines.. which use this one channel
	// then in the main() we will use multiple <-channel as code line to receive multiple values
	// but through the same channel

	// Lets create one channel
	channel := make(chan bool)

	// Now lets create multiple go-routines that take same amount of time to execute
	// Note: I want to create a race condition
	go greetWithTime_1(1, channel)
	go greetWithTime_2(1, channel)
	go greetWithTime_3(1, channel)

	// I will only output from the channel once
	// <- channel // ---- this code is commented ----

	// Here in this line 3 go-routines will send 'true' value. But we are outputting only one time.
	// What if happen is very simple, any of these 3 function calls.. whichever one of these functions
	// who get executed first (milli second difference) will output to the terminal. Then main-thread has
	// no code left to execute so it will end program, and YES! all other two function;s output will not be
	// shown in terminal as main-thread did not wait some more milliseconds for them to execute.
	// Also these 2 functions send the boolean 'true' to the channel but it was never outputted in the main()
	// So we need multiple redundant <-channel to make the main-thread wait for all these functions.
	//
	// Example: Let's say by some milli seconds greetWithTime_3() inputs to the channel first, Now main-thread
	// is just waiting there for someone to output something to the <-channel. When main-thread receives the
	// 'true' from greetWithTime_3() first.. it immiediately ends program and those 2 remaining functions
	// were not fast enough to input to the channel first so they are immidiately dumped as the program ends.
	// So their output never gets to the terminal and one function who wins (which can be anyone here) is able to
	// show its output in the terminal.
	//
	// So what is the solution ? Well just add more redundant <-channel in main(). Their amount should be 3 here
	// because i have 3 go-routine-functions here and only 3 times input is being send to my single channel.
	//
	// So amount of time input is being sent to the channel == amount of redundant <-channel
	//

	// Adding redundant lines to let every go-routine finish
	<-channel
	<-channel
	<-channel

	// Always remember the output will not come in the order in which the functions were called
	// but in a random order in which go-routines finish themselves.
	// Because go-routines are running separately from each other and main thread in background
	// so their output can have some random delay (milli seconds, seconds etc.) due to which their output
	// can come in un-ordered manner mostly.
}

// -- Channel pattern 1
//
//	What if you have multiple go-routine-functions ? What if you have multiple channels ?
//  How to make complex scenarios of channels and go-routines work ?
//	Well there are many ways and patterns (basically pattern of writing code and making a flow)
//	to use channels in different ways.
//
//	The purpose is the same: to make all go-routines finish themselves before main-thread ends program.
//	But some patterns and ways can help in some scenarios to make things work
//
