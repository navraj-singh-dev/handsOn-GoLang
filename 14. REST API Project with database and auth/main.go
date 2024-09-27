package main

import (
	"net/http"

	"e.com/events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// gives a pre-configured http server
	server := gin.Default()

	// --- Routes ---

	// return all stored events
	server.GET("/events", getEvents)

	// create a event
	server.POST("/events", createEvent)

	// start the server and listen on some port
	server.Run(":8080") // localhost:8080
}

// -- Handler Functions --

func getEvents(context *gin.Context) {
	// this slice contains all event struct instances that were created
	// remember slice is temporary as it is in RAM and not in a database
	// so it vanishes off when server restarts
	events := models.GetAllEvents()

	// return this slice as in JSON response
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// create a empty variable that stores event instance
	var event models.Event

	// .ShouldBindJSON() populates a &variable with the data incoming from the request,
	// So a user will send POST request with correct structure of 'event' struct instance data in request,
	// and this method will automatically fill the &variable with the data if it is in right structure
	// This method also supports 'struct tags' and i have define tags in event model struct also
	err := context.ShouldBindJSON(&event)

	// a error will come when a required field is missing in POST request's data
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "can't parse request data"})
		return
	}

	// populate other fields of event instance which are missing
	event.ID = 1
	event.UserID = 1

	// save the event instance to the temporary living slice
	event.Save()

	// send back the saved event instance as response
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})
}
