package routes

import (
	"net/http"

	"e.com/events-rest-api/models"
	"e.com/events-rest-api/utils"
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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "can't parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "cannot authenticate user..."})
		return
	}

	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "cannot authenticate user..."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"user logged-in successfully...", "JWT-token": token})
}
