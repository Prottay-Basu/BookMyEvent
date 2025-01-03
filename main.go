package main

import (
	"prottaybasu/book-my-event/db"
	"prottaybasu/book-my-event/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	// Sets up an HTTP server for us(already comes with recovery and logging middleware)
	server := gin.Default()

	routes.RegisterRoutes(server)

	// Starts the server and starts listening to the incoming request
	server.Run(":8080") //localhost:8080
}
