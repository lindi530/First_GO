package problem_model

import "time"

type ProblemSubmission struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	ProblemId   uint      `json:"problem_id"`
	Language    string    `json:"language"`
	Code        string    `json:"code"`
	State       string    `json:"state"`
	ExecTime    int       `json:"exec_time"`
	MemoryUsage int       `json:"memory_usage"`
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
	UserName    string    `json:"user_name" gorm:"-"`
}
