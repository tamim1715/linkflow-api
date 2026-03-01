package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppEnv = ""
var ServerPort string
var MongoURI string
var MongoDatabase string
var JWTSecret string
var JWTExpireHours string
var VerifyTokenURI string

func LoadEnv() {

	AppEnv = os.Getenv("APP_ENV")
	// Load .env only in development
	if AppEnv != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println(".env file not found")
		}
	}
	ServerPort = os.Getenv("SERVER_PORT")
	MongoURI = os.Getenv("MONGODB_URI")
	MongoDatabase = os.Getenv("MONGODB_DATABASE")
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTExpireHours = os.Getenv("JWT_EXPIRE_HOURS")
	VerifyTokenURI = os.Getenv("VERIFY_TOKEN_URI")
}
