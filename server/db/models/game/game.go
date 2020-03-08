package game

import "encoding/json"

type Game struct {
	Id          string `json:"id"`
	UserId      int32  `json:"userid,omitempty"`
	PointGained string `json:"pointsgained,omitempty"`
	WinStatus   string `json:"winstatus,omitempty"`
	GameType    string `json:"gametype,omitempty"`
	Created     string `json:"created,omitempty"`
	CreatedDay  string `json:"createdday,omitempty"`
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
