package user

import (
	"context"
	"errors"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
)

func Get(ctx context.Context, id int32) (*User, error) {
	u := &User{}
	filter := bson.D{{"id", id}}
	collection := mongo.GetClient().Database("db").Collection("users")
	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Delete(ctx context.Context, id int32) (bool, error) {
	filter := bson.D{{"id", id}}
	collection := mongo.GetClient().Database("db").Collection("users")
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	if res.DeletedCount == 1 {
		return true, nil
	}
	return false, errors.New("User was not deleted ")
}

func Update(ctx context.Context, user *User) (bool, error) {
	collection := mongo.GetClient().Database("db").Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return false, err
	}
	return true, nil
}

//func GetUsers(ctx context.Context, page int32, count int32) ([]User, error) {
//	collection := mongo.GetClient().Database("db").Collection("users")
//	lookupStage := bson.D{
//		{"$lookup", bson.D{
//			{"from", "games"},
//			{"localField", "id"},
//			{"foreignField", "userid"},
//			{"as", "count"},
//		}},
//	}
//	groupStage := bson.D{
//		{"$project", bson.D{
//			{"id", "$id"},
//			{"email", "$email"},
//			{"country", "$country"},
//			{"city", "$city"},
//			{"gender", "$gender"},
//			{"birthdate", "$birthdate"},
//			{"lastname", "$lastname"},
//			{"gamescount", bson.D{
//				{"$size", "$count"},
//			}},
//		}},
//	}
//	skipStage := bson.D{
//		{"$skip", int64((page - 1) * count)},
//	}
//	limitStage := bson.D{
//		{"$limit", count},
//	}
//
//	opts := options.Aggregate().SetAllowDiskUse(true)
//	cur, err := collection.Aggregate(ctx, mongo2.Pipeline{lookupStage, groupStage, skipStage, limitStage}, opts)
//	if err != nil {
//		return nil, err
//	}
//	users := []User{}
//	usr := User{}
//	for cur.Next(ctx) {
//		usr = User{}
//		err := cur.Decode(&usr)
//		if err != nil {
//			return nil, err
//		}
//
//		users = append(users, usr)
//	}
//	cur.Close(ctx)
//	return users, nil
//}

func GetUsers(ctx context.Context, page int32, count int32) ([]User, error) {
	collection := mongo.GetClient().Database("db").Collection("users")
	findOpt := options.Find()
	if page != 1 {
		findOpt.SetSkip(int64((page - 1) * count))
	}
	findOpt.SetSort(bson.D{{"id", 1}})
	findOpt.SetLimit(int64(count))
	cur, err := collection.Find(ctx, bson.D{{}}, findOpt)
	if err != nil {
		return nil, err
	}
	users := []User{}
	usr := User{}
	for cur.Next(ctx) {
		usr = User{}
		err := cur.Decode(&usr)
		if err != nil {
			return nil, err
		}

		users = append(users, usr)
	}
	cur.Close(ctx)
	collection = mongo.GetClient().Database("db").Collection("games")
	matchStage := bson.D{
		{"$match",
			bson.D{
				{
					"userid", bson.D{
						{"$gte", users[0].Id},
						{"$lte", users[len(users)-1].Id},
					},
				},
			},
		},
	}
	countStage := bson.D{
		{"$group", bson.D{
			{"_id", "$userid"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	cur, err = collection.Aggregate(ctx, mongo2.Pipeline{matchStage, countStage})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	sort.SliceStable(results, func(i, j int) bool {
		return results[i]["_id"].(int32) < results[j]["_id"].(int32)
	})
Main:
	for num, user := range users {
		for _, result := range results {
			if user.Id == result["_id"].(int32) {
				users[num].GamesCount = result["count"].(int32)
				continue Main
			}
		}
	}

	cur.Close(ctx)
	return users, nil
}
