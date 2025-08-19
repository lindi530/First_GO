package match_model

type MatchUser struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	Avatar       string `json:"avatar"`
	Rating       int    `json:"rating"`
	Level        string `json:"level"`
	Wins         int    `json:"wins"`
	TotalMatches int    `json:"total_matches"`
}
