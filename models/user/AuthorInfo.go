package user

type AuthorInfo struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}
