package problem_model

type RunResult struct {
	Passed   bool   `json:"passed"`
	Output   string `json:"output"`
	Expected string `json:"expected"`
	Error    string `json:"error"`
}
