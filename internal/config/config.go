package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppName        string
	AppEnv         string
	ServerPort     string
	MongoURI       string
	MongoDatabase  string
	JWTSecret      string
	JWTExpireHours string
}

func Load() *Config {

	// Load .env only in development
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println(".env file not found")
		}
	}

	return &Config{
		AppName:        getEnv("APP_NAME", "LinkFlow"),
		AppEnv:         getEnv("APP_ENV", "development"),
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		MongoURI:       getEnv("MONGODB_URI", ""),
		MongoDatabase:  getEnv("MONGODB_DATABASE", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpireHours: getEnv("JWT_EXPIRES_HOURS", "168"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
