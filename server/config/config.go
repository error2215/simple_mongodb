package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiPort   string
	MongoPort string
}

var GlobalConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}
	GlobalConfig = Config{
		ApiPort:   os.Getenv("API_PORT"),
		MongoPort: os.Getenv("MONGO_PORT"),
	}
}
