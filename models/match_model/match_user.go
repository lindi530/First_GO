package match_model

type MatchUser struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}
