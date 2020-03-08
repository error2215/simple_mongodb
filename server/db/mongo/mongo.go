package mongo

import (
	"context"
	"github.com/error2215/simple_mongodb/server/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client

func init() {
	var err error
	if os.Getenv("DOCKER_COMPOSE") == "true" {
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:" + config.GlobalConfig.MongoPort))
	} else {
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:" + config.GlobalConfig.MongoPort))
	}
	if err != nil {
		log.WithField("method", "server.db.mongo.init").Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.WithField("method", "server.db.mongo.init").Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.WithField("method", "server.db.mongo.init").Fatal(err)
	}
	log.Info("Connection to MongoDB finished. Address: " + config.GlobalConfig.MongoPort)
}

func GetClient() *mongo.Client {
	return client
}
