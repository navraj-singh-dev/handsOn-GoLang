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

// using close() here because this function is set to end at last
func greetWithTime_3(sec time.Duration, channel chan bool) {
	time.Sleep(sec * time.Second)
	fmt.Printf("\nGreetings _3 ! After %v seconds", sec)

	channel <- true

	close(channel)

}

func main() {

	channel := make(chan bool)

	go greetWithTime_1(1, channel)
	go greetWithTime_2(1, channel)
	// I need to consiously give more seconds to this function below so i can use close() in it,
	// otherwise this pattern cannot be used.
	go greetWithTime_3(3, channel) // giving 3 seconds so that this goRoutine ends at last

	// the special pattern is just to direcly iterate over the single channel
	for range channel { // this very line is automatically same as <-channel so we dont need write this

		// no code needed to write inside of the for-loop
		// we can omit _, _ := or _, value := or whatever is written.. because we don't want to use no index
		// or value here. We can directly write : for range channel.
		//
	}

}

// -- Channel pattern 3 (great choice if you know which go routine will end last)

// -- This pattern is used when you know which goRoutine will take the longest to complete
//	  otherwise dont use this.
//
//	Here a special pattern of for-loop is used and also a close(<channel>) function is used.
//
//	When a new channel is created it is Opened, but using close() we can permanently close that channel,
//  in all of the previous folders, i did not close any channels and did not use this close() function.
//
//	In this pattern there will be only 1 channel only used across all the multiple goRoutines just like pattern
//	one. But we will use special for-loop for channels.
//
//	This for-loop do not iterate over the slice of channels, instead it directly loop over our single channel.
//	It iterates infinite times over the same channel and never stops. Why? well because the channel is still opened
//	and special for-loop thinks i still output the data at some point, so main-thread is blocked for eternity
//	at some iteration of this for-loop. Due to this golang will print a error in the terminal but do not worry
//	all the go-routines will and whole code will run properly.. just the error will be printed in the terminal.
//
//	This is why a close(<channel>) is used in one of the goRoutines so that this for-loop know when to stops
//	iterating and release the main-thread. Otherwise be ready to see the error in terminal or console.
//
//	Iterating with for-loop over the channel directly will give the value inside of the channel, in this
//	example it is boolean, so it will give 'true' as it is what passed as input by the functions.
//
//	So mostly inside of the for-loop we don't need to write any code as <-channel. As it is already being
//  done by for-loop definition.
//
//
