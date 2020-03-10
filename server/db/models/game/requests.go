package game

import (
	"context"
	"github.com/error2215/go-convert"
	"github.com/error2215/simple_mongodb/server/db/models/mapSlice"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGamesGroupedByDateAndNumber(ctx context.Context) ([]GroupResult, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"createdday", "$createdday"},
				{"gametype", "$gametype"}}},
			{"gamesCount", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	secondStage := bson.D{
		{"$group", bson.D{
			{"_id", "$_id.createdday"},
			{"gametypes", bson.D{
				{"$push", bson.D{
					{"gametype", "$_id.gametype"},
					{"count", "$gamesCount"},
				}},
			}},
			{"count", bson.D{{"$sum", "$gamesCount"}}}},
		},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage, secondStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	res := []GroupResult{}
	for _, result := range results {
		groupRes := GroupResult{
			Day:        result["_id"].(string),
			GamesCount: result["count"].(int32),
		}
		for _, mapp := range result["gametypes"].(primitive.A) {
			groupRes.GameTypes = append(groupRes.GameTypes, GameType{
				Count:    convert.Int32(mapp.(primitive.M)["count"].(int32)),
				GameType: convert.Int32(mapp.(primitive.M)["gametype"].(string)),
			})
		}
		res = append(res, groupRes)

	}
	cur.Close(ctx)
	return res, nil
}

func GetRating(ctx context.Context, page int32, count int32) (mapSlice.MapSlice, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$userid"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	sortStage := bson.D{
		{"$sort", bson.D{
			{"count", -1},
		}},
	}
	skipStage := bson.D{
		{"$skip", int64((page - 1) * count)},
	}
	limitStage := bson.D{
		{"$limit", count},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage, sortStage, skipStage, limitStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	ms := mapSlice.MapSlice{}
	for _, result := range results {
		ms = append(ms, mapSlice.MapItem{Key: result["_id"].(int32), Value: result["count"].(int32)})
	}
	cur.Close(ctx)
	return ms, nil
}
