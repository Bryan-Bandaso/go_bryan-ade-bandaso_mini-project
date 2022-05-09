package dto

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
