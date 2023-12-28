package auth_test

import (
	context "context"
	"testing"
	"time"

	domain "github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/middleware"
	"github.com/akmalulginan/carjod-be/services/auth/usecase"
	"github.com/akmalulginan/carjod-be/services/user"
	"github.com/akmalulginan/carjod-be/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthUsecase_Login(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	passswordHash, _ := utils.HashPassword("123123123")

	expectedUser := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:          30,
		Email:        "udin@mail",
		Gender:       "male",
		PasswordHash: passswordHash,
	}

	data := domain.Auth{
		Email:    "udin@mail",
		Password: "123123123",
	}

	token := middleware.NewJWTAuthService().GenerateToken(expectedUser)

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(expectedUser, nil)

	result, err := authUsecase.Login(ctx, data)

	assert.NoError(t, err)
	assert.Equal(t, token, result)
}

func TestAuthUsecase_Login_Failed_EMAIL(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	data := domain.Auth{
		Email:    "udin@mails",
		Password: "123qweasds",
	}

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(domain.User{}, nil)

	_, err := authUsecase.Login(ctx, data)

	assert.Error(t, err)
}

func TestAuthUsecase_Login_Failed_PASSWORD(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	passswordHash, _ := utils.HashPassword("123123123")
	expectedUser := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:          30,
		Email:        "udin@mail",
		Gender:       "male",
		PasswordHash: passswordHash,
	}

	data := domain.Auth{
		Email:    "udin@mail",
		Password: "xxxxxx",
	}

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(expectedUser, nil)

	_, err := authUsecase.Login(ctx, data)

	assert.Error(t, err)
}

func TestAuthUsecase_Register(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	data := domain.Auth{
		Email:      "udin@mail",
		Password:   "xxxxxx",
		RePassword: "xxxxxx",
	}

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(domain.User{}, nil)
	mockUserRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	err := authUsecase.Register(ctx, data)

	assert.NoError(t, err)
}
func TestAuthUsecase_Register_Failed_EMAIL(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	userId := uuid.NewV4().String()

	passswordHash, _ := utils.HashPassword("123123123")
	expectedUser := domain.User{
		Model: domain.Model{Id: userId},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:          30,
		Email:        "udin@mail",
		Gender:       "male",
		PasswordHash: passswordHash,
	}

	data := domain.Auth{
		Email:    "udin@mail",
		Password: "xxxxxx",
	}

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(expectedUser, nil)

	err := authUsecase.Register(ctx, data)

	assert.Error(t, err)
}

func TestAuthUsecase_Register_Failed_PASSWORD(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockUserRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	data := domain.Auth{
		Email:      "udin@mail",
		Password:   "xxxxxx",
		RePassword: "zzzzzz",
	}

	mockUserRepo.EXPECT().FindByEmail(ctx, data.Email).Return(domain.User{}, nil)

	err := authUsecase.Register(ctx, data)

	assert.Error(t, err)
}
