package main

import (
	"log"

	"github.com/abcdataorg/sarmaaya-ticketing-backend/config"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/database"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/routes"
)

func main() {
	// Load configuration
	// config.LoadConfig()
	cfg := config.GetConfig()

	// Set server mode from configuration
	// gin.SetMode(config.AppConfig.Server.Mode)

	// Initialize database
	database.ConnectDatabase()

	// Set up router
	router := routes.SetupRouter()

	// Start server with dynamically set port
	// port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	if err := router.Run(cfg.PORT); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
