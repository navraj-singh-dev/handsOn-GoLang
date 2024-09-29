package middlewares

import (
	"net/http"

	"e.com/events-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	clientToken := context.Request.Header.Get("Authorization")

	if clientToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no JWT token provided..."})
		return
	}

	userId, err := utils.VerifyToken(clientToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "JWT token verification failed..."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
