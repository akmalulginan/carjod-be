package usecase

import (
	"context"
	"os"
	"time"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/brianvoe/gofakeit/v6"
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

func (u userUsecase) GenerateUser(ctx context.Context, qty int) error {
	gofakeit.Seed(0)
	defaultPassword, err := utils.HashPassword(os.Getenv("GENERATE_USER_PASSWORD"))
	if err != nil {
		return err
	}

	startDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Now().Location())
	endDate := time.Date(2006, 1, 1, 0, 0, 0, 0, time.Now().Location())

	users := make([]domain.User, 0)
	for i := 0; i < qty; i++ {
		user := domain.User{
			Name:         gofakeit.Name(),
			DateOfBirth:  utils.DateYYMMDD{Time: gofakeit.DateRange(startDate, endDate)},
			Email:        gofakeit.Email(),
			PasswordHash: defaultPassword,
			Gender:       gofakeit.RandomString([]string{"male", "female"}),
		}
		for j := 0; j < 2; j++ {
			hobby := gofakeit.RandomString([]string{"reading", "traveling", "cooking", "coding", "dancing", "gaming", "hiking", "sports"})
			user.Hobbies = append(user.Hobbies, hobby)
		}

		users = append(users, user)
	}

	u.userRepository.CreateBulk(ctx, &users)

	return nil
}
