package user

import (
	"context"
	"errors"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetUsers(ctx context.Context, page int32, count int32) ([]User, error) {
	collection := mongo.GetClient().Database("db").Collection("users")
	findOpt := options.Find()
	if page != 1 {
		findOpt.SetSkip(int64((page - 1) * count))
	}
	//findOpt.SetSort("id") //TODO check err
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
	return users, nil
}
