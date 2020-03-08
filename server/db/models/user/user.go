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
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type User struct {
	Id         int32  `json:"_id"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	GamesCount int32  `json:"games_count, omitempty"`
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

	for i, v := range mockUserStruct.Objects {
		v.Id = convert.Int32(i)
	}

	games = mockGameStruct.Objects

	allUserGames := []interface{}{}
	collection := mongo.GetClient().Database("db").Collection("games")
	counter := 0
	for _, localUser := range mockUserStruct.Objects {
		counter++
		userGames := getRandomGames(localUser.Id)
		localUser.GamesCount = int32(len(userGames))
		allUserGames = append(allUserGames, userGames...)
		if counter == 100 {
			_, err = collection.InsertMany(context.TODO(), allUserGames)
			if err != nil {
				log.Fatal(err)
			}
			allUserGames = []interface{}{}
			counter = 0
		}
	}

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
	currentCursor = 0
)

func getRandomGames(userId int32) []interface{} {
	num := rand.Intn(maxGamesCount-minGamesCount) + minGamesCount
	if currentCursor+num > usersCount-1 {
		currentCursor = 0
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(games), func(i, j int) { games[i], games[j] = games[j], games[i] })
	}
	slice := games[currentCursor : currentCursor+num]
	intSlice := []interface{}{}
	for num, curGame := range slice {
		curGame.UserId = userId
		curGame.Id = convert.String(userId) + "_" + convert.String(num)
		intSlice = append(intSlice, interface{}(curGame))
	}
	currentCursor = currentCursor + num
	logrus.Info("currentCursor ", currentCursor, " numberOfGames ", num, " UserId ", userId)
	return intSlice
}

type mockUser struct {
	Objects []User `json:"objects"`
}

type mockGame struct {
	Objects []game.Game `json:"objects"`
}
