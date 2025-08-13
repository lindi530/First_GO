package problem_model

type Example struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type ProblemTag struct {
	ID        int    `json:"id"`
	ProblemID uint   `json:"problem_id"`
	Tag       string `json:"tag"`
}

type ProblemExample struct {
	ID        int    `json:"id"`
	ProblemID uint   `json:"problem_id"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

type Problem struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Level       string `json:"level"`
	Description string `json:"description"`
	InputDesc   string `json:"input_desc"`
	OutputDesc  string `json:"output_desc"`
	TimeLimit   int    `json:"time_limit"`   // 单位秒
	MemoryLimit int    `json:"memory_limit"` // 单位 MB
	SubmitCount int    `json:"submit_count"`
	AcCount     int    `json:"ac_count"`

	Tags     []string  `json:"tags" gorm:"-"`
	Examples []Example `json:"examples" gorm:"-"`
}
