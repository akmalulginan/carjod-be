package usecase

import (
	"context"
	"time"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(ir domain.UserRepository) domain.UserUsecase {
	return userUsecase{userRepository: ir}
}

func (u userUsecase) GetById(ctx context.Context, userId string) (user domain.User, err error) {
	user, err = u.userRepository.FindById(ctx, userId)
	if err != nil {
		return user, err
	}

	loc, err := utils.GetLocaleJKT()
	if err != nil {
		return user, err
	}

	if !user.DateOfBirth.IsZero() {

		now := time.Now().In(loc)

		age := now.Year() - user.DateOfBirth.Year()

		if now.YearDay() < user.DateOfBirth.YearDay() {
			age--
		}

		user.Age = age
	}

	return user, nil
}

func (u userUsecase) Edit(ctx context.Context, user *domain.User) (err error) {
	return u.userRepository.Update(ctx, user)
}
