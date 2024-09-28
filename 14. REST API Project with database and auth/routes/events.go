package routes

import (
	"net/http"
	"strconv"

	"e.com/events-rest-api/models"
	"github.com/gin-gonic/gin"
)

// -- Handler Functions --

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch events, try again later..."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id..."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "cannot fetch event by its id..."})
		return
	}

	context.JSON(http.StatusOK , event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "can't parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event , try again later..."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})
}
