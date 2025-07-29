package problem_model

type Example struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type ProblemTag struct {
	ID        int    `json:"id"`
	ProblemID string `json:"problem_id"`
	Tag       string `json:"tag"`
}

type ProblemExample struct {
	ID        int    `json:"id"`
	ProblemID string `json:"problem_id"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

type Problem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Level       string `json:"level"`
	Description string `json:"description"`
	InputDesc   string `json:"input_desc"`
	OutputDesc  string `json:"output_desc"`

	Tags     []string  `json:"tags" gorm:"-"`
	Examples []Example `json:"examples" gorm:"-"`
}
