package problem_model

type CodeSubmission struct {
	Language  string `json:"language"`
	Code      string `json:"code"`
	ProblemID int64  `json:"problem_id"`
}
