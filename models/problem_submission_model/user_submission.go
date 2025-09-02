package problem_submission_model

import "time"

type UserSubmission struct {
	Id        int64     `json:"id"`
	ProblemId uint      `json:"problem_id"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
