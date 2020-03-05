package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	RESTPort  string
	GRPCPort  string
	MongoPort string
}

var GlobalConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}
	GlobalConfig = Config{
		RESTPort:  os.Getenv("REST_PORT"),
		GRPCPort:  os.Getenv("GRPC_PORT"),
		MongoPort: os.Getenv("MONGO_PORT"),
	}
}
