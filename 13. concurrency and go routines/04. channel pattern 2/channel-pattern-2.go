package main

import (
	"fmt"
	"time"
)

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

	// a slice holding 3 empty indexes
	// remember all indexes are empty and hold no value(channels) because make() is used
	channelSlice := make([]chan bool, 3)

	// Sending each channel from the slice as argument to the functions
	// Now all channels in the slicee are send as argument one by one to each go-routine-function
	// Each function will give input to the channel.. inside of their function body code
	channelSlice[0] = make(chan bool) // placing a channel as value at index 0
	go greetWithTime_1(1, channelSlice[0])
	channelSlice[1] = make(chan bool)
	go greetWithTime_2(1, channelSlice[1])
	channelSlice[2] = make(chan bool)
	go greetWithTime_3(1, channelSlice[2])

	// Using for-loop to iterate over the list of channels to check for outputs
	for _, channel := range channelSlice {
		<-channel
	}

	// One clarity i wanna give is that during for-loop dont think that for-loop will iterate quickly and
	// fast over the slice of channel and end instantly. This could lead to confusions.
	//
	// For every iteration of for-loop the main-thread will be blocked or waiting on the <-channel
	// until some output is recieved from it. If the output of <-channel comes after let's say 6 seconds
	// then in 1st iteration main-thread will be blocked/waiting for 6 second as it is expecting output
	// from <-channel. Then main-thread will be waiting/blocked for 2nd and 3rd iteration of for-loop
	// also. So overall this for-loop will take high amount of time and the time taken by this for-loop
	// totally depends on when <-channel output's something.
	//
}

// -- Channel pattern 2 (good to use, other patterns are also available but this still can be used)
//
//	Create multiple channels and put them in a slice. Each channel in the slice will be mapped to the
//  a single go-routine-function.
//
//	Each go-routine-function will input data to its own channel it is mapped to. Mapping is not a new code
//	or concept.. instead i am simply saying that just create as many channels as their are go-routine-functions
//	and put them in a slice. Then one channel from slice will be used by one go-routine-function and that is it.
//
//	Use a for-loop in the main() and iterate over the channels to get the output.
//
