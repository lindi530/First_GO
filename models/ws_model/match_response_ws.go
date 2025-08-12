package ws_model

import "GO1/models/match_model"

type MatchResponse struct {
	Type      string                 `json:"type"`
	To        int64                  `json:"to"`
	RoomID    string                 `json:"roomId"`
	ProblemID uint                   `json:"problemId"`
	Opponent  *match_model.MatchUser `json:"opponent"`
}
