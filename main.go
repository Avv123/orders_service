package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"orderservice/config"
	"orderservice/consumers"
	"orderservice/repository"
	"orderservice/routes"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if config.DB == nil {
		log.Fatal("Database client is nil after connection")
	}
	repository.InitOrderCollection()
	router := gin.Default()
	go consumers.StartEmailConsumer()
	//routes.SetRoutes(router)

	routes.SetupRoutes(router)
	//sockets.Initsocket()
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
