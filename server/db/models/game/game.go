package game

import "encoding/json"

type Game struct {
	Id          string `json:"id"`
	UserId      int32  `json:"user_id,omitempty"`
	PointGained string `json:"points_gained,omitempty"`
	WinStatus   string `json:"win_status,omitempty"`
	GameType    string `json:"game_type,omitempty"`
	Created     string `json:"created,omitempty"`
	CreatedDay  string `json:"created_day,omitempty"`
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
