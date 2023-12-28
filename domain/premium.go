package domain

import (
	"context"

	"github.com/akmalulginan/carjod-be/utils"
)

type Premium struct {
	UserId    string        `json:"user_id"`
	Type      utils.Premium `json:"type"`
	IsSuccess bool          `json:"is_success"`
}

type PremiumUsecase interface {
	Upgrade(ctx context.Context, data Premium) error
	Webhook(ctx context.Context, data Premium) error
}
