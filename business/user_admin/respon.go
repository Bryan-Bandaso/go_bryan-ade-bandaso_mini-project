package user

import "github.com/ydhnwb/golang_heroku/entity"

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
