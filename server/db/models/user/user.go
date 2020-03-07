package user

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/error2215/go-convert"
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
	Id        int32   `json:"_id"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Gender    string  `json:"gender"`
	BirthDate string  `json:"birth_date"`
	GamesIds  []int32 `json:"games_ids"`
}

func SliceToJson(users ...*User) ([]byte, error) {
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
	users = []interface{}{}
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
	usersInt := []interface{}{}
	for i, v := range mockUserStruct.Objects {
		v.Id = convert.Int32(i)
		usersInt = append(usersInt, interface{}(v))
	}
	//jsonTmp := map[string]interface{}{}
	//err = json.Unmarshal(usersBytes, &jsonTmp)
	//if err != nil{
	//	logrus.Error(err)
	//	os.Exit(1)
	//}
	//jsonArr := jsonTmp["objects"].([]interface{})
	//for num, intUser := range jsonArr{
	//	localUser := User{}
	//	usrMarsh,err := json.Marshal(intUser)
	//	if err != nil{
	//		logrus.Error(err)
	//		os.Exit(1)
	//	}
	//	err = json.Unmarshal(usrMarsh, &localUser)
	//	if err != nil{
	//		logrus.Error(err)
	//		os.Exit(1)
	//	}
	//	localUser.Id = convert.Int32(num)
	//	users = append(users, interface{}(localUser))
	//}
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//										Games select
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//games = []game.Game{}
	//gamesBytes := []byte{}
	//data = make([]byte, 64)
	//
	//file, err = os.Open("source/games.json")
	//if err != nil{
	//	logrus.Error(err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//for {
	//	n, err := file.Read(data)
	//	if err == io.EOF{
	//		break
	//	}
	//	gamesBytes = append(gamesBytes, data[:n]...)
	//}
	//jsonTmp = map[string]interface{}{}
	//err = json.Unmarshal(gamesBytes, &jsonTmp)
	//if err != nil{
	//	logrus.Error(err)
	//	os.Exit(1)
	//}
	//jsonArr = jsonTmp["objects"].([]interface{})
	//for _, intGame := range jsonArr{
	//	localGame := game.Game{}
	//	gameMarsh,err := json.Marshal(intGame)
	//	if err != nil{
	//		logrus.Error(err)
	//		os.Exit(1)
	//	}
	//	err = json.Unmarshal(gameMarsh, &localGame)
	//	if err != nil{
	//		logrus.Error(err)
	//		os.Exit(1)
	//	}
	//	games = append(games, localGame)
	//}
	//logrus.Info(len(games))
	//logrus.Info(games[50000])
	//

	collection := mongo.GetClient().Database("db").Collection("users")
	insertManyResult, err := collection.InsertMany(context.TODO(), usersInt)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("Inserted users len: " + convert.String(len(insertManyResult.InsertedIDs)))

	//allUserGames := []interface{}{}
	//collection = mongo.GetClient().Database("db").Collection("games")
	//counter := 0
	//for _, localUser := range users{
	//	counter++
	//	userGames := getRandomGames(localUser.(User).Id)
	//	allUserGames = append(allUserGames, userGames...)
	//	if counter == 30 {
	//		insertManyResult, err = collection.InsertMany(context.TODO(), allUserGames)
	//		logrus.Info("sent")
	//		allUserGames = []interface{}{}
	//		counter = 0
	//	}
	//}
	//logrus.Info("mda")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//logrus.Info("Inserted gae len: " + convert.String(len(insertManyResult.InsertedIDs)))
}

var games []game.Game
var users []interface{}
var min = 5000
var max = 15000
var currentCursor = 0

func getRandomGames(userId int32) []interface{} {
	num := rand.Intn(max-min) + min
	if currentCursor+num > 99999 {
		currentCursor = 0
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(games), func(i, j int) { games[i], games[j] = games[j], games[i] })
	}
	slice := games[currentCursor : currentCursor+num]
	logrus.Info(len(slice))
	intSlice := []interface{}{}
	for num, curGame := range slice {
		curGame.UserId = userId
		curGame.Id = convert.String(userId) + "_" + convert.String(num)
		intSlice = append(intSlice, interface{}(curGame))
	}
	currentCursor = currentCursor + num
	logrus.Info("cur ", currentCursor, " num ", num, " id ", userId)
	return intSlice
}

type mockUser struct {
	Objects []User `json:"objects"`
}
