package game

type Game struct {
	Id          string `json:"_id"`
	UserId      int32  `json:"user_id"`
	PointGained string `json:"point_gained"`
	WinStatus   string `json:"win_status"`
	GameType    string `json:"game_type"`
	Created     string `json:"created"`
}
