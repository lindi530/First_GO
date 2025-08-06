package ws_model

type MatchSuccess struct {
	RoomID    string `json:"roomId"`
	ProblemID uint   `json:"problemId"`
	Opponent  string `json:"opponent"`
}
