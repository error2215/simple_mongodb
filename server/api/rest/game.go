package rest

import (
	"encoding/json"
	"github.com/error2215/simple_mongodb/server/db/models"
	"github.com/error2215/simple_mongodb/server/db/models/game"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) GetGamesByNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("GetGamesByNumberHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	//we can run alternative method that works with aggregation
	//u, err := game.GetGamesGroupedByNumberAggregation(r.Context())
	u, err := game.GetGamesGroupedByNumber(r.Context())
	if err != nil {
		log.Errorf("GetGamesByNumberHandler/game.GetGamesGroupedByNumber() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	encoded, err := json.Marshal(u)
	if err != nil {
		log.Errorf("GetGamesByNumberHandler/json.Marshal() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", encoded).ToString())
}

func (s *Server) GetGamesByDateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("GetGamesByDateHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	//we can run alternative method that works with aggregation
	//u, err := game.GetGamesGroupedByDateAggregation(r.Context())
	u, err := game.GetGamesGroupedByDate(r.Context())
	if err != nil {
		log.Errorf("GetGamesByDateHandler/game.GetGamesGroupedByDate() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	encoded, err := json.Marshal(u)
	if err != nil {
		log.Errorf("GetGamesByDateHandler/json.Marshal() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", encoded).ToString())
}
