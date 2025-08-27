package problem_model

type BasicInfo struct {
	Title      string   `json:"title"`
	Difficulty string   `json:"difficulty"`
	Tags       []string `json:"tags"`
}

type Description struct {
	Description  string `json:"description"`
	InputFormat  string `json:"inputFormat"`
	OutputFormat string `json:"outputFormat"`
	Hint         string `json:"hint"`
}

type Constraint struct {
	TimeLimit      int `json:"timeLimit"`   // 毫秒
	MemoryLimit    int `json:"memoryLimit"` // MB
	MaxSubmissions int `json:"maxSubmissions"`
}
type Constraints struct {
	Cpp    Constraint `json:"cpp"`
	Python Constraint `json:"python"`
	Go     Constraint `json:"go"`
}

type TestCases struct {
	InputContent  string `json:"inputContent"`
	InputFile     string `json:"inputFile"`
	OutputContent string `json:"outputContent"`
	OutputFile    string `json:"outputFile"`
	IsPublic      bool   `json:"isPublic"`
}

type UploadProblem struct {
	BasicInfo   BasicInfo   `json:"basicInfo"`
	Description Description `json:"description"`
	TestCases   []TestCases `json:"testCases"`
	Constraints Constraints `json:"constraints"`
}

func (pd *UploadProblem) MapProblem() *Problem {

	examples := make([]Example, len(pd.TestCases))
	for i, tc := range pd.TestCases {
		examples[i] = Example{
			Input:    tc.InputContent,
			Output:   tc.OutputContent,
			IsPublic: tc.IsPublic,
		}
		if i == 0 {
			examples[i].IsPublic = true
		}
	}

	problem := Problem{
		Title:       pd.BasicInfo.Title,
		Level:       pd.BasicInfo.Difficulty,
		Description: pd.Description.Description,
		InputDesc:   pd.Description.InputFormat,
		OutputDesc:  pd.Description.OutputFormat,
		SubmitCount: 0,
		AcCount:     0,

		Constraints: pd.Constraints,
		Tags:        pd.BasicInfo.Tags,
		Examples:    examples,
	}

	return &problem
}
