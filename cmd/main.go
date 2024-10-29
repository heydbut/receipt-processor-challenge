package main

import (
	"log"

	"receiptprocessor/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	handlers.SetupRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
