package problem_model

type ProblemList struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Level       string   `json:"level"`
	SubmitCount int      `json:"submit_count"`
	AcCount     int      `json:"ac_count"`
	Accepted    bool     `json:"accepted"`
	Tags        []string `json:"tags" gorm:"-"`
	PassRate    float64  `json:"pass_rate"`
}
