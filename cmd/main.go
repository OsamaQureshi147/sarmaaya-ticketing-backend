package main

import (
	"log"

	"github.com/abcdataorg/sarmaaya-ticketing-backend/config"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/database"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/models"
	"github.com/abcdataorg/sarmaaya-ticketing-backend/routes"
)

func init() {
	// Initialize database
	database.ConnectDatabase()

	// Sync models
	// Auto-migrate all models
	for _, model := range models.RegisterModels() {
		if err := database.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model: %v", err)
		}
	}

	log.Println("All models migrated successfully.")
}

func main() {
	// Load configuration
	cfg := config.GetEnvConfig()

	// Set up router
	router := routes.SetupRouter()

	// Start server with dynamically set port
	// port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	if err := router.Run(cfg.PORT); err != nil {
		log.Fatalf("Server failed to hello start: %v", err)
	}
}
