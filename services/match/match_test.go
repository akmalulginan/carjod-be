package match_test

import (
	context "context"
	"testing"
	"time"

	domain "github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/services/match"
	"github.com/akmalulginan/carjod-be/services/match/usecase"
	"github.com/akmalulginan/carjod-be/services/tx"
	"github.com/akmalulginan/carjod-be/services/user"
	"github.com/akmalulginan/carjod-be/utils"
	gomock "github.com/golang/mock/gomock"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestMatchUsecase_GetCandidate(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMatchRepo := match.NewMockMatchRepository(ctrl)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	mockTxCoordinator := tx.NewMockTxCoordinator(ctrl)

	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo, mockTxCoordinator)

	expectedUser := domain.User{
		Model: domain.Model{Id: uuid.NewV4().String()},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:     30,
		Email:   "udin@mail",
		Gender:  "male",
		Hobbies: pq.StringArray{"gaming", "cooking"},
	}

	expectedCandidate := domain.User{
		Model: domain.Model{Id: uuid.NewV4().String()},
		Name:  "Sarah",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Email:  "sarah@mail",
		Gender: "female",
	}

	expectedMatches := []domain.Match{
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
	}

	mockUserRepo.EXPECT().FindById(ctx, expectedUser.Id).Return(expectedUser, nil)
	mockMatchRepo.EXPECT().FindByUserId(ctx, expectedUser.Id, false, false).Return(expectedMatches, nil)
	mockUserRepo.EXPECT().FindCandidate(ctx, expectedUser, gomock.Any()).Return(expectedCandidate, nil)

	result, err := matchUseCase.GetCandidate(ctx, expectedUser.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedCandidate, result)
}

func TestMatchUsecase_GetMatches(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMatchRepo := match.NewMockMatchRepository(ctrl)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	mockTxCoordinator := tx.NewMockTxCoordinator(ctrl)

	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo, mockTxCoordinator)

	userId := uuid.NewV4().String()

	expectedMatches := []domain.Match{
		{
			UserId:       userId,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       userId,
			TargetUserId: uuid.NewV4().String(),
		},
	}

	expectedUserMatches := []domain.User{
		{
			Model: domain.Model{Id: expectedMatches[0].Id},
			Name:  "Sarah",
			DateOfBirth: utils.DateYYMMDD{
				Time: time.Date(1995, 10, 9, 0, 0, 0, 0, time.Local),
			},
			Age:     30,
			Email:   "sarah@mail",
			Gender:  "female",
			Hobbies: pq.StringArray{"gaming", "cooking"},
		},
		{
			Model: domain.Model{Id: expectedMatches[0].Id},
			Name:  "Wati",
			DateOfBirth: utils.DateYYMMDD{
				Time: time.Date(1993, 11, 7, 0, 0, 0, 0, time.Local),
			},
			Age:     30,
			Email:   "sati@mail",
			Gender:  "female",
			Hobbies: pq.StringArray{"gaming", "sport"},
		},
	}

	mockMatchRepo.EXPECT().FindByUserId(ctx, userId, true, false).Return(expectedMatches, nil)
	mockUserRepo.EXPECT().FindByIds(ctx, gomock.Any()).Return(expectedUserMatches, nil)

	result, err := matchUseCase.GetMatches(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserMatches, result)
}

func TestMatchUsecase_Action(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMatchRepo := match.NewMockMatchRepository(ctrl)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	mockTxCoordinator := tx.NewMockTxCoordinator(ctrl)

	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo, mockTxCoordinator)

	expectedUser := domain.User{
		Model: domain.Model{Id: uuid.NewV4().String()},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:     30,
		Email:   "udin@mail",
		Gender:  "male",
		Hobbies: pq.StringArray{"gaming", "cooking"},
	}

	expectedMatches := []domain.Match{
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
	}

	data := domain.Match{
		UserId:       expectedUser.Id,
		TargetUserId: uuid.NewV4().String(),
		IsLike:       true,
	}

	mockUserRepo.EXPECT().FindById(ctx, expectedUser.Id).Return(expectedUser, nil)

	mockMatchRepo.EXPECT().FindByUserId(ctx, expectedUser.Id, false, true).Return(expectedMatches, nil)
	mockMatchRepo.EXPECT().FindLiked(ctx, data.TargetUserId, expectedUser.Id).Return(domain.Match{}, nil)
	mockMatchRepo.EXPECT().Create(ctx, &data).Return(nil)

	mockTxCoordinator.EXPECT().Begin(ctx).Return(ctx, nil)
	mockTxCoordinator.EXPECT().Rollback(ctx).Return(nil)
	mockTxCoordinator.EXPECT().Commit(ctx).Return(nil)

	err := matchUseCase.Action(ctx, &data)

	assert.NoError(t, err)
}

func TestMatchUsecase_Action_LIMIT_REACHED(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMatchRepo := match.NewMockMatchRepository(ctrl)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	mockTxCoordinator := tx.NewMockTxCoordinator(ctrl)

	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo, mockTxCoordinator)

	expectedUser := domain.User{
		Model: domain.Model{Id: uuid.NewV4().String()},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:     30,
		Email:   "udin@mail",
		Gender:  "male",
		Hobbies: pq.StringArray{"gaming", "cooking"},
	}

	expectedMatches := []domain.Match{
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
	}

	data := domain.Match{
		UserId:       expectedUser.Id,
		TargetUserId: uuid.NewV4().String(),
		IsLike:       true,
	}

	mockUserRepo.EXPECT().FindById(ctx, expectedUser.Id).Return(expectedUser, nil)

	mockMatchRepo.EXPECT().FindByUserId(ctx, expectedUser.Id, false, true).Return(expectedMatches, nil)

	err := matchUseCase.Action(ctx, &data)

	assert.Error(t, err)
}
func TestMatchUsecase_Action_PREMIUM_UNLIMITED(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMatchRepo := match.NewMockMatchRepository(ctrl)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	mockTxCoordinator := tx.NewMockTxCoordinator(ctrl)

	matchUseCase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo, mockTxCoordinator)

	expectedUser := domain.User{
		Model: domain.Model{Id: uuid.NewV4().String()},
		Name:  "Udin",
		DateOfBirth: utils.DateYYMMDD{
			Time: time.Date(1993, 10, 7, 0, 0, 0, 0, time.Local),
		},
		Age:            30,
		Email:          "udin@mail",
		Gender:         "male",
		Hobbies:        pq.StringArray{"gaming", "cooking"},
		PremiumActicve: true,
		PremiumSwipe:   true,
	}

	expectedMatches := []domain.Match{
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
		{
			UserId:       expectedUser.Id,
			TargetUserId: uuid.NewV4().String(),
		},
	}

	data := domain.Match{
		UserId:       expectedUser.Id,
		TargetUserId: uuid.NewV4().String(),
		IsLike:       true,
	}

	mockUserRepo.EXPECT().FindById(ctx, expectedUser.Id).Return(expectedUser, nil)

	mockMatchRepo.EXPECT().FindByUserId(ctx, expectedUser.Id, false, true).Return(expectedMatches, nil)
	mockMatchRepo.EXPECT().FindLiked(ctx, data.TargetUserId, expectedUser.Id).Return(domain.Match{}, nil)
	mockMatchRepo.EXPECT().Create(ctx, &data).Return(nil)

	mockTxCoordinator.EXPECT().Begin(ctx).Return(ctx, nil)
	mockTxCoordinator.EXPECT().Rollback(ctx).Return(nil)
	mockTxCoordinator.EXPECT().Commit(ctx).Return(nil)

	err := matchUseCase.Action(ctx, &data)

	assert.NoError(t, err)
}
