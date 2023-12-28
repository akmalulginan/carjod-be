package user_test

import (
	context "context"
	"testing"
	"time"

	domain "github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/services/user"
	"github.com/akmalulginan/carjod-be/services/user/usecase"
	"github.com/akmalulginan/carjod-be/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_GetById(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	userUseCase := usecase.NewUserUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	expectedUser := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:    30,
		Email:  "udin@mail",
		Gender: "male",
	}

	mockUserRepo.EXPECT().FindById(ctx, userId).Return(expectedUser, nil)

	resultUser, err := userUseCase.GetById(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, resultUser)

}

func TestUserUseCase_Edit(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	userUseCase := usecase.NewUserUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	user := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:    30,
		Email:  "udin@mail",
		Gender: "male",
	}

	mockUserRepo.EXPECT().Update(ctx, &user).Return(nil)

	err := userUseCase.Edit(ctx, &user)

	assert.NoError(t, err)
}
