package api

import (
	"net/http"
)

type API interface {
	Start()

	GetUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)

	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	DeleteUsersHandler(w http.ResponseWriter, r *http.Request)
	CreateUsersHandler(w http.ResponseWriter, r *http.Request)
	UpdateUsersHandler(w http.ResponseWriter, r *http.Request)
}
