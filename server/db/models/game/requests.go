package game

import (
	"context"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

// Method with aggregation // average 11 sec
func GetGamesGroupedByNumberAggregation(ctx context.Context) (map[string]int32, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$gametype"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	res := make(map[string]int32)
	for _, result := range results {
		res[result["_id"].(string)] = result["count"].(int32)
	}
	cur.Close(ctx)
	return res, nil
}

// Method with goroutines // average 3 sec
func GetGamesGroupedByNumber(ctx context.Context) (map[string]int32, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$gametype"},
		}},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	res := make(map[string]int32)
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	for _, result := range results {
		wg.Add(1)
		go func(id interface{}, group *sync.WaitGroup, res map[string]int32, mutex *sync.RWMutex) {
			cur, _ := collection.CountDocuments(ctx, bson.D{{"gametype", id}})
			mutex.Lock()
			res[id.(string)] = int32(cur)
			mutex.Unlock()
			group.Done()
		}(result["_id"], &wg, res, &mutex)

	}
	wg.Wait()
	cur.Close(ctx)
	return res, nil
}

// Method with aggregation // average 15 sec
func GetGamesGroupedByDateAggregation(ctx context.Context) (map[string]int32, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$createdday"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	res := make(map[string]int32)
	for _, result := range results {
		res[result["_id"].(string)] = result["count"].(int32)
	}
	cur.Close(ctx)
	return res, nil
}

// Method with goroutines // average 4 sec
func GetGamesGroupedByDate(ctx context.Context) (map[string]int32, error) {
	collection := mongo.GetClient().Database("db").Collection("games")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$createdday"},
		}},
	}
	opts := options.Aggregate().SetAllowDiskUse(true)
	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{groupStage}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	res := make(map[string]int32)
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	for _, result := range results {
		wg.Add(1)
		go func(id interface{}, group *sync.WaitGroup, res map[string]int32, mutex *sync.RWMutex) {
			cur, _ := collection.CountDocuments(ctx, bson.D{{"createdday", id}})
			mutex.Lock()
			res[id.(string)] = int32(cur)
			mutex.Unlock()
			group.Done()
		}(result["_id"], &wg, res, &mutex)

	}
	wg.Wait()
	cur.Close(ctx)
	return res, nil
}
