package usecase

import (
	"context"
	"errors"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/middleware"
	"github.com/akmalulginan/carjod-be/utils"
)

type authUsecase struct {
	userRepository domain.UserRepository
}

func NewAuthUsecase(ur domain.UserRepository) domain.AuthUsecase {
	return authUsecase{userRepository: ur}
}

func (u authUsecase) Login(ctx context.Context, data domain.Auth) (token string, err error) {

	user, err := u.userRepository.FindByEmail(ctx, data.Email)
	if err != nil || user.Id == "" {
		return "", errors.New("invalid email")
	}

	passwordHash, err := utils.HashPassword(data.Password)
	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(passwordHash, data.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token = middleware.NewJWTAuthService().GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u authUsecase) Register(ctx context.Context, data domain.Auth) (err error) {
	user, err := u.userRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if user.Id != "" {
		return errors.New("email has been registered")
	}

	if data.Password != data.RePassword {
		return errors.New("password doesn't match")
	}

	passwordHash, err := utils.HashPassword(data.Password)
	if err != nil {
		return err
	}

	user = domain.User{
		Email:        data.Email,
		PasswordHash: passwordHash,
	}

	err = u.userRepository.Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
