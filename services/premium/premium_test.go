package premium_test

import (
	context "context"
	"testing"
	"time"

	domain "github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/services/premium/usecase"
	"github.com/akmalulginan/carjod-be/services/user"
	"github.com/akmalulginan/carjod-be/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestPremiumUsecase_Upgrade(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	premiumUseCase := usecase.NewPremiumUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	expectedUserFind := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:    30,
		Email:  "udin@mail",
		Gender: "male",
	}

	expectedUserUpdate := expectedUserFind
	expectedUserUpdate.PremiumSwipe = true

	data := domain.Premium{
		UserId: userId,
		Type:   utils.PremiumSwipe,
	}

	mockUserRepo.EXPECT().FindById(ctx, userId).Return(expectedUserFind, nil)
	mockUserRepo.EXPECT().Update(ctx, &expectedUserUpdate).Return(nil)

	err := premiumUseCase.Upgrade(ctx, data)

	assert.NoError(t, err)
}

func TestPremiumUsecase_Upgrade_FAILED(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	premiumUseCase := usecase.NewPremiumUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	expectedUserFind := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:            30,
		Email:          "udin@mail",
		Gender:         "male",
		PremiumSwipe:   true,
		PremiumActicve: true,
	}

	data := domain.Premium{
		UserId: userId,
		Type:   utils.PremiumVerified,
	}

	mockUserRepo.EXPECT().FindById(ctx, userId).Return(expectedUserFind, nil)

	err := premiumUseCase.Upgrade(ctx, data)

	assert.Error(t, err)
}

func TestPremiumUsecase_Webhook(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	premiumUseCase := usecase.NewPremiumUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	expectedUserFind := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Email:        "udin@mail",
		Gender:       "male",
		PremiumSwipe: true,
	}

	expectedUserUpdate := expectedUserFind
	expectedUserUpdate.PremiumActicve = true

	data := domain.Premium{
		UserId:    userId,
		IsSuccess: true,
	}

	mockUserRepo.EXPECT().FindById(ctx, userId).Return(expectedUserFind, nil)
	mockUserRepo.EXPECT().Update(ctx, &expectedUserUpdate).Return(nil)

	err := premiumUseCase.Webhook(ctx, data)

	assert.NoError(t, err)
}
