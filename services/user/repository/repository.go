package repository

import (
	"context"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/lib/pq"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return userRepository{db: db}
}

func (r userRepository) FindById(ctx context.Context, id string) (user domain.User, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	return user, err
}

func (r userRepository) FindCandidate(ctx context.Context, user domain.User, matchIds []string) (candidate domain.User, err error) {
	db := r.db.
		Where("id != ?", user.Id).
		Where("gender != ?", user.Gender).
		Where("hobbies && ?", pq.Array(user.Hobbies)).
		Order("RANDOM()")
	if len(matchIds) > 0 {
		db = db.Where("id NOT IN ?", matchIds)
	}
	err = db.First(&candidate).Error
	return candidate, err
}

func (r userRepository) FindByIds(ctx context.Context, ids []string) (users []domain.User, err error) {
	err = r.db.WithContext(ctx).Where("id in ?", ids).Find(&users).Error
	return users, err
}

func (r userRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r userRepository) Create(ctx context.Context, user *domain.User) (err error) {
	err = r.db.WithContext(ctx).Create(user).Error
	return err
}

func (r userRepository) CreateBulk(ctx context.Context, users *[]domain.User) (err error) {
	err = r.db.WithContext(ctx).Create(&users).Error
	return err
}

func (r userRepository) Update(ctx context.Context, user *domain.User) (err error) {
	err = r.db.WithContext(ctx).Where("id = ?", user.Id).Updates(user).Error
	return err
}
