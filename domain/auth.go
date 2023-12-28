package domain

import "context"

type Auth struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type AuthUsecase interface {
	Register(ctx context.Context, data Auth) (err error)
	Login(ctx context.Context, data Auth) (token string, err error)
}
