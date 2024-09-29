package routes

import (
	"e.com/events-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// -- event table route --
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// -- group routes (protected routes)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// protected events table routes
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// protected registrations table route
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// -- user table routes --
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
