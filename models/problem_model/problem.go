package problem_model

type Example struct {
	Input    string `json:"input"`
	Output   string `json:"output"`
	IsPublic bool   `json:"is_public"`
}

type ProblemTag struct {
	ProblemID uint   `json:"problem_id"`
	Tag       string `json:"tag"`
}

type ProblemExample struct {
	ProblemID uint   `json:"problem_id"`
	Input     string `json:"input"`
	Output    string `json:"output"`
	IsPublic  bool   `json:"is_public"`
}

type Problem struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Level       string `json:"level"`
	Description string `json:"description"`
	InputDesc   string `json:"input_desc"`
	OutputDesc  string `json:"output_desc"`

	Constraints Constraints `json:"constraints" gorm:"-"`
	SubmitCount int         `json:"submit_count" gorm:"-"`
	AcCount     int         `json:"ac_count" gorm:"-"`

	Tags     []string  `json:"tags" gorm:"-"`
	Examples []Example `json:"examples" gorm:"-"`
}
