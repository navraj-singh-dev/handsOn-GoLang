package routes

import (
	"net/http"
	"strconv"

	"e.com/events-rest-api/models"
	"e.com/events-rest-api/utils"
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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch event by its id..."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	clientToken := context.Request.Header.Get("Authorization")

	if clientToken == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "no JWT token provided..."})
		return
	}

	userId, err := utils.VerifyToken(clientToken)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token verification failed..."})
		return
	}

	var event models.Event

	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "can't parse request data"})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event , try again later..."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id..."})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch event by id, try again later..."})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse request's data..."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot update event , try again later..."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully..."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id..."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch event by id, try again later..."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot delete event by id, try again later..."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully..."})
}
