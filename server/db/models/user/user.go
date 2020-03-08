package user

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/error2215/go-convert"
	"github.com/error2215/simple_mongodb/server/config"
	"github.com/error2215/simple_mongodb/server/db/models/game"
	"github.com/error2215/simple_mongodb/server/db/mongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

type User struct {
	Id         int32  `json:"id"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
	Gender     string `json:"gender,omitempty"`
	BirthDate  string `json:"birthdate,omitempty"`
	GamesCount int32  `json:"gamescount,omitempty"`
}

func SliceToJson(users ...User) ([]byte, error) {
	encoded, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func (s *User) ToJson() ([]byte, error) {
	encoded, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

//TODO Refactor method for better performance
func GenerateMock() {
	err := mongo.GetClient().Database("db").Drop(context.TODO())
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	mod := mongo2.IndexModel{
		Keys: bson.M{
			"gametype": 1, // index in ascending order
		}, Options: options.Index().SetUnique(false),
	}
	_, err = mongo.GetClient().Database("db").Collection("games").Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	mod = mongo2.IndexModel{
		Keys: bson.M{
			"createdday": 1, // index in ascending order
		}, Options: options.Index().SetUnique(false),
	}
	_, err = mongo.GetClient().Database("db").Collection("games").Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	//										Users select
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	usersBytes := []byte{}
	file, err := os.Open("source/users_go.json")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		usersBytes = append(usersBytes, data[:n]...)
	}
	r := bytes.NewReader(usersBytes)
	decoder := json.NewDecoder(r)
	var mockUserStruct *mockUser
	err = decoder.Decode(&mockUserStruct)
	if err != nil {
		logrus.Error(err)
		return
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//										Games select
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	gamesBytes := []byte{}
	data = make([]byte, 64)

	file, err = os.Open("source/games.json")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		gamesBytes = append(gamesBytes, data[:n]...)
	}
	r = bytes.NewReader(gamesBytes)
	decoder = json.NewDecoder(r)
	var mockGameStruct *mockGame
	err = decoder.Decode(&mockGameStruct)
	if err != nil {
		logrus.Error(err)
		return
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	for i := range mockUserStruct.Objects {
		mockUserStruct.Objects[i].Id = convert.Int32(i)
		if i+1 == usersCount {
			break
		}
	}
	games = mockGameStruct.Objects

	collection := mongo.GetClient().Database("db").Collection("games")
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	waitNumber := 800
	for i, localUser := range mockUserStruct.Objects {
		if i == waitNumber {
			waitNumber += 800
			wg.Wait()
		}
		wg.Add(1)
		go func(id int32, mutex *sync.RWMutex) {
			userGames := getRandomGames(id, mutex)
			mockUserStruct.Objects[i].GamesCount = int32(len(userGames))
			_, err = collection.InsertMany(context.TODO(), userGames)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(localUser.Id, &mutex)
		if i+1 == usersCount {
			break
		}
	}
	wg.Wait()
	usersInt := []interface{}{}
	for i, v := range mockUserStruct.Objects {
		usersInt = append(usersInt, interface{}(v))
		if i+1 == usersCount {
			break
		}
	}
	collection = mongo.GetClient().Database("db").Collection("users")
	insertManyResult, err := collection.InsertMany(context.TODO(), usersInt)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("Inserted users len: " + convert.String(len(insertManyResult.InsertedIDs)))
}

var (
	games         []game.Game
	minGamesCount = config.GlobalConfig.MockMinGamesCount
	maxGamesCount = config.GlobalConfig.MockMaxGamesCount
	usersCount    = config.GlobalConfig.MockUsersCount
)

func getRandomGames(userId int32, mutex *sync.RWMutex) []interface{} {
	num := rand.Intn(maxGamesCount+1-minGamesCount) + minGamesCount
	mutex.Lock()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(games), func(i, j int) { games[i], games[j] = games[j], games[i] })
	mutex.Unlock()
	slice := games[0:num]
	intSlice := []interface{}{}
	for num, curGame := range slice {
		curGame.UserId = userId
		curGame.Id = convert.String(userId) + "_" + convert.String(num)
		for i, num := range curGame.Created {
			if string(num) == " " {
				curGame.CreatedDay = curGame.Created[:i]
				break
			}
		}
		intSlice = append(intSlice, interface{}(curGame))
	}
	logrus.Info("numberOfGames: ", num, " UserId: ", userId, " Goroutine num:", runtime.NumGoroutine())
	return intSlice
}

type mockUser struct {
	Objects []User `json:"objects"`
}

type mockGame struct {
	Objects []game.Game `json:"objects"`
}

//9 mln 100 sec goroutines 500 no sync limit
//7 mln 100 sec goroutines 500 sync limit
//6.8 mln 100 sec goroutines 500 sync no limit
