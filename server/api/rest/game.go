package rest

import (
	"errors"
	"github.com/error2215/go-convert"
	"github.com/error2215/simple_mongodb/server/db/models"
	"github.com/error2215/simple_mongodb/server/db/models/game"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) GetRatingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Errorf("GetRatingHandler/ParseForm() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	page := r.FormValue("page")
	count := r.FormValue("count")
	if (page == "") || (count == "") {
		log.Errorf("GetRatingHandler/FormValue() err: %v", errors.New("Missing required params"))
		_, _ = w.Write(models.Response(1, errors.New("Missing required params").Error(), nil).ToString())
		return
	}
	u, err := game.GetRating(r.Context(), convert.Int32(page), convert.Int32(count))
	if err != nil {
		log.Errorf("GetRatingHandler/game.GetRating() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	jsonData, err := game.SliceToJson(u...)
	if err != nil {
		log.Errorf("GetRatingHandler/game.SliceToJson() err: %v", err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(models.Response(0, "", jsonData).ToString())
}
