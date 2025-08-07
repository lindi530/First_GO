package ws_model

type MatchSuccess struct {
	Type      string `json:"type"`
	To        int64  `json:"to"`
	RoomID    string `json:"roomId"`
	ProblemID uint   `json:"problemId"`
	Opponent  string `json:"opponent"`
}
