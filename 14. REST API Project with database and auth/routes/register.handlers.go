package routes

import (
	"net/http"
	"strconv"

	"e.com/events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id..."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot not fetch event..."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot not register user for event..."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered user for event..."})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id..."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot not cancel registration..."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registration canceled..."})
}
