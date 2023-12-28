package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
)

type premiumUsecase struct {
	userRepository domain.UserRepository
}

func NewPremiumUsecase(ur domain.UserRepository) domain.PremiumUsecase {
	return premiumUsecase{userRepository: ur}
}

func (u premiumUsecase) Upgrade(ctx context.Context, data domain.Premium) (err error) {
	user, err := u.userRepository.FindById(ctx, data.UserId)
	if err != nil || user.Id == "" {
		return errors.New("invalid user")
	}

	if user.PremiumActicve {
		premiumActive := utils.PremiumSwipe
		if user.PremiumVerified {
			premiumActive = utils.PremiumVerified
		}
		return fmt.Errorf("you are already %s", premiumActive)
	}

	user.PremiumSwipe = data.Type == utils.PremiumSwipe
	user.PremiumVerified = data.Type == utils.PremiumVerified

	return u.userRepository.Update(ctx, &user)
}

func (u premiumUsecase) Webhook(ctx context.Context, data domain.Premium) (err error) {

	if !data.IsSuccess {
		return errors.New("failed to upgrade")
	}

	user, err := u.userRepository.FindById(ctx, data.UserId)
	if err != nil || user.Id == "" {
		return errors.New("invalid user")
	}

	user.PremiumActicve = data.IsSuccess

	return u.userRepository.Update(ctx, &user)
}
