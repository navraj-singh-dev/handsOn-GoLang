package main

import (
	"e.com/events-rest-api/db"
	"e.com/events-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// create/open database connection
	db.InitDB()

	// gives a pre-configured http server
	server := gin.Default()

	// outsource routes in another package
	routes.RegisterRoutes(server)

	// start the server and listen on some port
	server.Run(":8080") // localhost:8080
}
