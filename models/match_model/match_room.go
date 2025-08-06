package match_model

import "time"

type Room struct {
	ID        int64     `json:"id"`
	RoomID    string    `json:"room_id"`
	User1ID   int64     `json:"user_1_id"`
	User2ID   int64     `json:"user2_id"`
	ProblemID uint      `json:"problem_id"`
	StartTime time.Time `json:"start_time"`
	Status    string    `json:"status"`
	WinnerID  int64     `json:"winner_id"`
}
