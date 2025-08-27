package problem_model

type CodeSubmission struct {
	Language  string `json:"language"`
	Code      string `json:"code"`
	ProblemID uint   `json:"problem_id"`
}
