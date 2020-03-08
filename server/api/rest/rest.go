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
		//r.Delete("/", s.DeleteUsersHandler) // remove users
		//r.Put("/", s.UpdateUsersHandler)    // add new users
		r.Get("/", s.GetUsersHandler) // get users
		//r.Post("/", s.CreateUsersHandler)   // update user's data
	})

	r.Route("/games", func(r chi.Router) {
		//r.Delete("/", s.DeleteUsersHandler) // remove games
		//r.Put("/", s.UpdateUsersHandler)    // add new games
		//r.Get("/", s.GetUsersHandler)       // get games
		r.Get("/rating", s.GetRatingHandler) // get ratings
		//r.Get("/{number}", s.GetGamesByNumberHandler)       // get games grouped by numbers of games
		//r.Get("/{day}", s.GetGamesByDateHandler)       // get games grouped by date
		//r.Post("/", s.CreateUsersHandler)   // update games's data
	})

	log.Info("Api Server started on port: " + config.GlobalConfig.ApiPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.ApiPort, r)
	if err != nil {
		log.WithField("ERROR", "Cannot start API Server").Fatal(err)
	}
}
