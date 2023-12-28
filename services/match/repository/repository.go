package repository

import (
	"context"
	"time"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"

	"gorm.io/gorm"
)

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) domain.MatchRepository {
	return matchRepository{db: db}
}

func (r matchRepository) FindByUserId(ctx context.Context, userId string, isMatch, isToday bool) (matches []domain.Match, err error) {
	db := r.db.WithContext(ctx).Where("user_id = ?", userId)
	if isToday {
		loc, _ := utils.GetLocaleJKT()
		startOfDay := time.Now().In(loc).Truncate(24 * time.Hour)
		db = db.Where("created_at >= ?", startOfDay)
	}
	if isMatch {
		db = db.Where("is_match = ?", isMatch)
	}
	err = db.Find(&matches).Error
	return matches, err
}
func (r matchRepository) FindLiked(ctx context.Context, userId, targetId string) (match domain.Match, err error) {
	err = r.db.WithContext(ctx).Where("user_id = ? AND target_user_id = ? AND is_like = ?", userId, targetId, true).Find(&match).Error
	return match, err
}

func (r matchRepository) Create(ctx context.Context, match *domain.Match) (err error) {
	err = r.db.WithContext(ctx).Create(match).Error
	return err
}

func (r matchRepository) Update(ctx context.Context, match *domain.Match) (err error) {
	err = r.db.WithContext(ctx).Where("id = ?", match.Id).Updates(match).Error
	return err
}
