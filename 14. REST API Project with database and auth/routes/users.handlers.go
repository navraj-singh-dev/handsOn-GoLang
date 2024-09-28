package routes

import (
	"net/http"

	"e.com/events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "can't parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "can't save user..."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"user created successfully..."})
}
