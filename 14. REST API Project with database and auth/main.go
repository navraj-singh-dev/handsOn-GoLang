package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gives a pre-configured http server
	server := gin.Default()

	// --- Routes ---
	// give a endpoint string then the handler function
	server.GET("/events", getEvents)

	// start the server and listen on some port
	// make sure this line of code is at last in main()
	server.Run(":8080") // localhost:8080
}

// -- Handler Functions --
func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}

// -- Some explanation on multiple things

//	A handler function should have *gin.Context as parameter. Because gin automatically passes this
//	as argument to every single handler. Why? Well just because of this *gin.Context we can send responses
//	back from our server to the client for the client's request. So this "context" argument is very
//	important. Gin expects you to define this parameter in your handler functions so dont forget it.
//
//	context.JSON() function/method is used to send back a response. It takes a status code and then
//	any kind of data and sends it back to the client as response in JSON format. Here i am sending a
//	map[string]any as response which will be converted to a JSON object automatically by .JSON()
//
//	gin.H{} is nothing but a alias (custom type) to map[string]any which i am sending back as a response.
//	So instead of writing map[string]any{<key>:<value>}.. which is bit long to write, we can use this alias
//	gin.H{} and simply send back a response in a map, which ofcourse will be converted to the JSON object
//	by .JSON() methods behind the scenes.
//
//	Also while sending http requests to this route a very interesting error i found,
//  that is if in the main()
//  server.Run(":8080") line of code is defined/written.. above any endpoint function like:
//  server.GET() | server.POST() | server.PATCH() etc. or any other http endpoint methods..
//	the server will not be able to listen on 'any endpoint at all' and give '404 not found'
//	So make sure server.Run(":8080") is at the most last of the main() function and it is the
//	last part of the code other wise no endpoint method like .GET()m .POST() etc. will work.
//
