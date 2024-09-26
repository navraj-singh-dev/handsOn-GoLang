package main

import (
	"errors"
	"fmt"
	"time"
)

func returnError(errorChan chan error) {
	errorChan <- errors.New("function failed with errors")
}

func timeConsuming(sec time.Duration, errorChan chan error, successChan chan bool, giveError bool) {

	time.Sleep(sec * time.Second)
	fmt.Printf("\nHi am i am time consuming function!\nHaving Errors? : %v\n", giveError)

	if giveError {
		returnError(errorChan)
		return
	}
	successChan <- true
}

func main() {
	errorChannel := make(chan error)
	successChannel := make(chan bool)

	go timeConsuming(2, errorChannel, successChannel, true)

	// -- define case using channels
	select {
		// which ever case of these two output's first wins and select statement ends immidiately
		// without even checking other cases
		// winning case get's its code executed.. while losing cases do not
	case err := <-errorChannel:
		if err != nil {
			fmt.Println(err)
		}
	case successValue := <-successChannel:
		fmt.Println("Function successfully executed!", successValue)
	}
}

// -- Channel Pattern 3 : Error Channel + Select Statement (Perfect Combo)

//	The syntax is : select{} + case <:variable> := <-channel_name: <code>
//
//	Well if you are in a situation where you must handle many different channels which are fulfilling
//	different purposes but you want to only select one channel who outputs first and then not care about
//	other channels.. then select statement if the solution.
//
//	In my case i defined the problem in previous folder, so select statement completely solves that.
//	The main-thread is no longer blocked and stuck and ends the program perfectly.
//
//	A select statement only works with channels, in its cases we define the all channels we want to,
//	the select statement works such that.. whatever channel-case output's first out of all the cases wins
//	and its code is executed. Then select statement exits without caring about all other cases.
//
//	In short in select statement the channel (defined as a case) which outputs first wins and all other
//	lose, the winner gets its code excutesd (the code written in the case) and losers code is never
//	executed. So only one channel out of all the channels in the select statement wins, that is exactly what
//	we want.
//
//  The select statement or any other pattern i discussed can be wrapped inside of the for-loop easily
//  if the scenario demands such need, other this way is also fine.
//
//  With this concurrency in golang and goRoutines ENDS!!
//	