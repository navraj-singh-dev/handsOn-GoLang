package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gives a pre-configured http server
	server := gin.Default()

	// --- Routes ---
	server.GET("/events", getEvents)

	// start the server and listen on some port
	server.Run(":8080") // localhost:8080
}

// -- Handler Functions --
func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
