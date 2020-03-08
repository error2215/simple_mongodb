package game

import "encoding/json"

type Game struct {
	Id          string `json:"_id"`
	UserId      int32  `json:"user_id"`
	PointGained string `json:"point_gained"`
	WinStatus   string `json:"win_status"`
	GameType    string `json:"game_type"`
	Created     string `json:"created"`
}

func SliceToJson(users ...Game) ([]byte, error) {
	encoded, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func (s *Game) ToJson() ([]byte, error) {
	encoded, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
