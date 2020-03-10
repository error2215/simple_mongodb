package rest

import (
	"github.com/error2215/simple_mongodb/server/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type Server struct {
}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {
		r.Delete("/{id}", s.DeleteUserHandler) // remove user
		r.Put("/", s.UpdateUserHandler)        // add new user
		r.Get("/{id}", s.GetUserHandler)       // get user by id
		r.Post("/", s.CreateUserHandler)       // update user's data
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", s.GetUsersHandler) // get users
	})

	r.Route("/games", func(r chi.Router) {
		r.Get("/rating", s.GetRatingHandler)              // get ratings
		r.Get("/group", s.GetGamesByDateAndNumberHandler) // get games grouped by date and number
	})

	log.Info("Api Server started on port: " + config.GlobalConfig.ApiPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.ApiPort, r)
	if err != nil {
		log.WithField("ERROR", "Cannot start API Server").Fatal(err)
	}
}
