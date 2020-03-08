package game

import (
	"context"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"sort"
)

func GetRating(ctx context.Context, page int32, count int32) ([]Game, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	collection = mongo.GetClient().Database("db").Collection("games")
	countStage := bson.D{
		{"$group", bson.D{
			{"_id", "$userid"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{countStage})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	sort.SliceStable(results, func(i, j int) bool {
		return results[i]["count"].(int32) > results[j]["count"].(int32)
	})
	for num, result := range results {
		logrus.Info(num, ": ", result)
	}

	cur.Close(ctx)
	return []Game{}, nil
}
