package config

import (
	"github.com/error2215/go-convert"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiPort      string
	MongoPort    string
	GenerateMock bool
}

var GlobalConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}
	GlobalConfig = Config{
		ApiPort:      os.Getenv("API_PORT"),
		MongoPort:    os.Getenv("MONGO_PORT"),
		GenerateMock: convert.Bool(os.Getenv("GENERATE_MOCK")),
	}
}
