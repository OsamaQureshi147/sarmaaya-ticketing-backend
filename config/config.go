package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT                 string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	SarmaayaClickUpToken string
	ClickUpApiUrl        string
}

func GetEnvConfig() *EnvConfig {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &EnvConfig{
		PORT:                 os.Getenv("PORT"),
		DBHost:               os.Getenv("DB_HOST"),
		DBPort:               os.Getenv("DB_PORT"),
		DBUser:               os.Getenv("DB_USER"),
		DBPassword:           os.Getenv("DB_PASSWORD"),
		DBName:               os.Getenv("DB_NAME"),
		SarmaayaClickUpToken: os.Getenv("SARMAAYA_BOT_API_TOKEN"),
		ClickUpApiUrl:        os.Getenv("CLICKUP_API_URL"),
	}
}
