package domain

import (
	"context"

	"github.com/akmalulginan/carjod-be/utils"
	"github.com/lib/pq"
)

type User struct {
	Model
	Name            string           `json:"name,omitempty"`
	DateOfBirth     utils.DateYYMMDD `json:"date_of_birth,omitempty"`
	Age             int              `json:"age,omitempty" gorm:"-"`
	Email           string           `json:"email,omitempty"`
	PasswordHash    string           `json:"-" gorm:"column:password_hash"`
	Gender          string           `json:"gender,omitempty"`
	Hobbies         pq.StringArray   `json:"hobbies,omitempty" gorm:"type:text[]"`
	PremiumSwipe    bool             `json:"premium_swipe"`
	PremiumVerified bool             `json:"premium_verified"`
	PremiumActicve  bool             `json:"premium_acticve"`
}

func (*User) TableName() string {
	return "user"
}

type UserUsecase interface {
	GetById(ctx context.Context, userId string) (user User, err error)
	Edit(ctx context.Context, user *User) (err error)
}

type UserRepository interface {
	FindById(ctx context.Context, id string) (user User, err error)
	FindByIds(ctx context.Context, ids []string) (users []User, err error)
	FindByEmail(ctx context.Context, email string) (user User, err error)
	FindCandidate(ctx context.Context, user User, matchIds []string) (candidate User, err error)
	Create(ctx context.Context, user *User) (err error)
	CreateBulk(ctx context.Context, users *[]User) (err error)
	Update(ctx context.Context, user *User) (err error)
}
