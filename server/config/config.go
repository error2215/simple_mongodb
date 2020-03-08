package config

import (
	"github.com/error2215/go-convert"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiPort                  string
	MongoPort                string
	GenerateMock             bool
	MockMinGamesCount        int
	MockMaxGamesCount        int
	MockGamesInsertBatchSize int
	MockUsersCount           int
}

var GlobalConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}
	GlobalConfig = Config{
		ApiPort:                  os.Getenv("API_PORT"),
		MongoPort:                os.Getenv("MONGO_PORT"),
		GenerateMock:             convert.Bool(os.Getenv("GENERATE_MOCK")),
		MockMaxGamesCount:        convert.Int(os.Getenv("MOCK_MAX_GAMES_COUNT")),
		MockMinGamesCount:        convert.Int(os.Getenv("MOCK_MIN_GAMES_COUNT")),
		MockGamesInsertBatchSize: convert.Int(os.Getenv("MOCK_GAMES_INSERT_BATCH_SIZE")),
		MockUsersCount:           convert.Int(os.Getenv("MOCK_USERS_COUNT")),
	}
}
