package main

import (
	"fmt"
	"time"
)

// add a parameter to accept a channel, as this is a go-routine function
func greetWithTime(sec time.Duration, channel chan bool) {

	// all this code is work, you can do any work you want normally
	time.Sleep(sec * time.Second)
	fmt.Printf("\nGreetings! After %v seconds", sec)

	// now in the end as work is completed, give some value as input of to channel
	// i just sent 'true' booleans as value to the channel as a input
	// now this value is in the channel
	// the main-thread in the main() will receive it
	channel <- true
}

func main() {

	// channel that communicates with bool data-type
	channel := make(chan bool)

	// main-thread Off Loads this task to a go-routine-thread and goes on next lines
	go greetWithTime(3, channel)

	<-channel // Take the output from the channel, it takes 3 seconds ofcourse
	// Here the main-thread stopped, now it must wait for this channel to output something.
	// channel will not output anything unless the go-routine-thread gets finished in the
	// 'greetWithTime()' function.

	// When the go-routine-thread will execute all the lines of code in 'greetWithTime()':
	// time.Sleep(sec * time.Second)
	// fmt.Printf("\nGreetings! After %v seconds", sec)

	// Then in the last line of 'greetWithTime()':
	// channel <- true
	
	// go-routine-thread will give the boolean value 'true' as input to the channel in 'greetWithTime()'.
	// only then.. in the main().. code line "<-channel".. will output that value(the 'true' boolean) 
	// in the main().. so now main-thread that was waiting for 3 seconds can take it
	// as output here in this line.

	// Now the main-thread as no line of code below the '<-channel' to execute and therefore it ends the
	// program. But our goal got accomplished which was to make main-thread always wait for go-routines
	// to end and only then end the program.

}

// -- Channels Keyword

//	All the functions which use "go" keyword or to say "go-routine-threads" must use a channel.
//	They must accept a channel in argument/parameter. Also in the code of the function body developer
//  must give a input value to channel of some sort. Then in the main() function the output of the this
//  channel must be received (taken out). For input and output to channels a "<-" operator is used.
//  This is the common flow of how channels + functions + go keyword are used together.
//
//	The main-thread never waits for any go routine, so we make it wait by using channels. The main thread
//	in the main() is blocked until a channel outputs something. That output will come through our go routine
//  function. As long as our function do not give a input to the channel in its function body.. this input
//	to the channel cannot be received as a output in the main(). So this way the main-thread waits.
//	Hence all our go routine functions will always have the honor to output their results without getting
//	betrayed the main-thread.
//
//	A channel is nothing but a communication device through which go-routines can communicate.
//  A channel can take any value as input and give it as output. For example as a go-routine-function
//	cannot return any value, like if a go-function needs to return error.. it cannot.. because we used
//	'go' keyworded function cannot. So the only way these type of functions can return any value is
//	channels. So it does not mean if it is a error or a calculated value or a object or a struct or
//	int or float what ever it is that you want this function to return.. it must be through that channel.
//	
//	So in short, main-thread runs the main() and it has the power to end the program by itself, it is the
//	leader-thread and all go-routine-threads are the worker-threads. By default main-thread do not wait for
//	any go-routine and ends the program, but channels are used to make the main-thread wait.
//	
