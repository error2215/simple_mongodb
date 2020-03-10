package rest

import (
	"encoding/json"
	"errors"
	"github.com/error2215/go-convert"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/error2215/simple_mongodb/server/db/models"
	"github.com/error2215/simple_mongodb/server/db/models/user"
)

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("GetUserHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	id := strings.Split(r.URL.String(), "/")[2] // 0 -> "", 1 -> "users", 2 -> {id}
	u, err := user.Get(r.Context(), convert.Int32(id))
	if err != nil {
		log.Errorf("GetUserHandler/user.Get() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	jsonData, err := u.ToJson()
	if err != nil {
		log.Errorf("GetUserHandler/user.SliceToJson() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", jsonData).ToString())
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("DeleteUserHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	id := strings.Split(r.URL.String(), "/")[2] // 0 -> "", 1 -> "users", 2 -> {id}
	ok, err := user.Delete(r.Context(), convert.Int32(id))
	if err != nil {
		log.Errorf("DeleteUserHandler/user.Delete() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	if ok != true {
		err := errors.New("User was not deleted due to unknown issues ")
		log.Errorf("DeleteUserHandler/user.Delete() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", nil).ToString())
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		log.Errorf("CreateUserHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var usr *user.User
	err = decoder.Decode(&usr)
	if err != nil {
		log.Errorf("CreateUserHandler/decoder.Decode() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	ok, err := user.Update(r.Context(), usr)
	if err != nil {
		log.Errorf("CreateUserHandler/user.Update() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	if ok != true {
		err := errors.New("User was not created due to unknown issues ")
		log.Errorf("CreateUserHandler/user.Update() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", nil).ToString())
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		log.Errorf("UpdateUserHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var usr *user.User
	err = decoder.Decode(&usr)
	if err != nil {
		log.Errorf("UpdateUserHandler/decoder.Decode() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	ok, err := user.Update(r.Context(), usr)
	if err != nil {
		log.Errorf("UpdateUserHandler/user.Update() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	if ok != true {
		err := errors.New("User was not updated due to unknown issues ")
		log.Errorf("UpdateUserHandler/user.Update() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", nil).ToString())
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("GetUsersHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	page := r.FormValue("page")
	count := r.FormValue("count")
	if (page == "") || (count == "") {
		log.Errorf("GetUsersHandler/FormValue() err: %v", errors.New("Missing required params"))
		_, _ = w.Write(models.Response(1, errors.New("Missing required params").Error(), nil).ToString())
		return
	}
	u, err := user.GetUsers(r.Context(), convert.Int32(page), convert.Int32(count))
	if err != nil {
		log.Errorf("GetUsersHandler/user.GetUsers() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	jsonData, err := user.SliceToJson(u...)
	if err != nil {
		log.Errorf("GetUsersHandler/user.SliceToJson() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", jsonData).ToString())
}

func (s *Server) DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {

}
