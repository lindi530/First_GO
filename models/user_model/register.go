package user_model

type ParamRegister struct {
	Name       string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,password"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
}
