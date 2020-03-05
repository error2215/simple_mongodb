package api

import "github.com/error2215/simple_mongodb/server/db/models/user"

type API interface {
	Start()

	GetUser(id int32) (user user.User)
	DeleteUser(id int32) (ok bool)
	CreateUser(user user.User) (ok bool)
	UpdateUser(user user.User) (ok bool)

	GetUsers(ids []int32) (users []user.User)
	DeleteUsers(ids []int32) (ok bool)
	CreateUsers(users []user.User) (ok bool)
	UpdateUsers(users []user.User) (ok bool)
}
