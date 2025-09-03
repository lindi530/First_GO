package post_model

type PageInfo struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"` // form 使用 ShouldBindQuery
}
