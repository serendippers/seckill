package responce

import "seckill/core/model"

type UserResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
