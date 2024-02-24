package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfigStruct struct {
	Port      string
	JWTSecret string
}

var AppConfig AppConfigStruct

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Warning: .env file not found")
	}

	AppConfig = AppConfigStruct{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET_KEY", "your_jwt_secret_key"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
