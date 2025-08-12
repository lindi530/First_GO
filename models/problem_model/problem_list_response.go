package problem_model

type ProblemListResponse struct {
	ID       uint     `json:"id"`
	Title    string   `json:"title"`
	Level    string   `json:"level"`
	Tags     []string `json:"tags" gorm:"-"`
	PassRate float64  `json:"pass_rate" gorm:"-"`
}
