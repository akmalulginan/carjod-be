package domain

import "context"

type Match struct {
	Model
	UserId       string `json:"user_id"`
	TargetUserId string `json:"target_user_id"`
	IsLike       bool   `json:"is_like"`
	IsMatch      bool   `json:"is_match"`
	User         *User  `gorm:"foreignKey:user_id"`
	TargetUser   *User  `gorm:"foreignKey:target_user_id"`
}

func (*Match) TableName() string {
	return "match"
}

type MatchUsecase interface {
	GetCandidate(ctx context.Context, userId string) (User, error)
	GetMatches(ctx context.Context, userId string) ([]User, error)
	Action(ctx context.Context, data *Match) error
}

type MatchRepository interface {
	FindByUserId(ctx context.Context, userId string, isMatch, isToday bool) ([]Match, error)
	FindLiked(ctx context.Context, userId, targetId string) (Match, error)
	Create(ctx context.Context, data *Match) error
	Update(ctx context.Context, data *Match) error
}
