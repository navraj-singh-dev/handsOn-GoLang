package main

import (
	"errors"
	"fmt"
	"time"
)

// A function that always return a error
// This function can never be a goRoutine because it returns something..
// But a goRoutine function can never return something..
// So we need to return errors also using Channels..

// -- This function is a example of a function that returns something..
// -- But now we need to rewrite it, so that it returns the error through a channel..
// -- otherwise it cannot be goRoutine function
// func returnError() error {
// 	fmt.Println("ERROR:")
// 	return errors.New("I am a error!")
// }

// Function that return error with a error channel
func returnError(errorChan chan error) {
	fmt.Println("ERROR:")
	fmt.Println(errors.New("i am a error"))
	errorChan <- errors.New("i am a error")
}

// a time consuming function that can give error on demand by passing boolean as argument
// if i pass true it will give, if i pass false i will not give error..
func timeConsuming(sec time.Duration, errorChan chan error, successChan chan bool, giveError bool) {

	time.Sleep(sec * time.Second)
	fmt.Printf("\nHi am i am time consuming function! %v\n", giveError)

	// give error if giveError is true
	// other wise this channel will never output anything
	if giveError {
		// putting error in this channel through a helper function
		returnError(errorChan)
		// now as error has occured i can return out of this very function
		return // a single empty return can be used in goRoutine just to exit out of this function
	}

	// if no error is send through errorChan then valChan will give value to tell function successfully worked
	successChan <- true
}

func main() {

	// create a channel.. but just for tramitting errors
	errorChannel := make(chan error)

	// create a channel but just for transmitting values other than errors
	successChannel := make(chan bool)

	// Send errors this time, errorChan outputs error, but not the successChan.. it never outputs anything.
	// This time this function fails and do not run successfully and returns a error through the errorchan
	// This can be compared with any other normal functions that actually return error with 'return' keyword
	// go timeConsuming(2, errorChannel, successChannel, true)

	// Don't send errors this time, errorChan never outputs anything, but the successChan outputs 'true' as boolean.
	// This time this function do not fail and run successfully and returns a 'true' boolean through the successChan
	// This can be compared with any other normal functions that actually return a success value with 'return' keyword
	go timeConsuming(2, errorChannel, successChannel, false)

	// wait for outputs from errorChan and valChan

	// here is a major problem when-even errorChannels are in the picture,
	// but that is dicussed below.
	<-errorChannel // main-thread will only go to next line if errorChannel outputs something otherwise its stuck here
	<-successChannel

}


// -- Channel Pattern 3 : Error Channel (needs a select statement, never use error Channels without select statement)

//	A good function (of any type) mostly returns two things when it is called:
//		A success-Value and nil-Error if it successfully runs and we use - value, err := pattern to get them.
//		Otherwise if it fails.. then it returns a nil as Success-Value and some error object as error which we will handle 
//	
//	But as these functions are converted to goRoutines.. they never return anything. They simply cannot.
//	So even if some error occurs in this function and it fails or it successfully runs we always have to  
//	use some channels to transmit those values out of the function as we cannot use 'return' keyword to 
//	return these error or success values anymore.
//	
//	To do this we create simple channel one with the name errorChannel and one with the name successChannel. 
//	Names can be anything but they must define their purpose. errorChannel just to send errors and successChannel
//	just to send success values out of the goRoutine function.
//	
//	Now the problem arises. In the main() main-thread will need to listen on both channels the error and also
//	the success one. But now the problem is what if the function executes successfully and error never comes and
//	therefore errorChannel will never output anything. This will make main thread to be stuck for eternity
//	on the <-errorChannel and it will never output anything and program will crash as main-thread is stuck.
//	Or what if the function fails and outputs error on <-errorChannel but now the <-successChannel stucks 
//	the main-thread because if error came then <-successChannel will never output anything only <-errorChannel
//	will output. So both ways main-thread is stuck.
//	
//	So the problem is we need both channels and also we need to listen on both channel with main-thread in main()
//	but how do we un-stuck main-thread ? and make such a logic that if error comes then <-successChannel do not
//	needs to be checked only <-errorChannel needs to be checked and if function successfully runs then only
//	<-successChannel needs to be checked and not the <-errorChannel. But right now both are being checked and 
//	both channels can never output at the same time. 
//	
//	For this we have select statement which is specially for channels only. So we will use that to handle 
//	this problem.
//	
//	This folder is just a PART 1 and do not provides whole solution, PART 2 in which i use select statement
//	is in next folder.