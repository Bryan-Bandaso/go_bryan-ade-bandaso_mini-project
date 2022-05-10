package user

import "project-art-museum/entity"

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `gorm:"type:varchar(100)" json:"-"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}
