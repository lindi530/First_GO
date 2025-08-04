package problem_model

type ProblemConstraint struct {
	ID          int    `json:"id"`
	ProblemID   uint   `json:"problem_id"`
	TimeLimit   int    `json:"time_limit"`
	MemoryLimit int    `json:"memory_limit"`
	Language    string `json:"language"`
}
