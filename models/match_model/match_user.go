package match_model

type MatchUser struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Rank     int    `json:"rank"`
}
