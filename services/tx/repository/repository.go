package repository

import (
	"context"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"gorm.io/gorm"
)

type txCoordinator struct {
	db *gorm.DB
}

func NewTxCoordinator(db *gorm.DB) domain.TxCoordinator {
	return &txCoordinator{db: db}
}

func (c *txCoordinator) Begin(ctx context.Context) (context.Context, error) {
	tx := c.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return context.WithValue(ctx, utils.KeyCoordinatorTx, tx), nil
}

func (c *txCoordinator) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(utils.KeyCoordinatorTx).(*gorm.DB)
	if !ok {
		return nil
	}
	return tx.Commit().Error
}

func (c *txCoordinator) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(utils.KeyCoordinatorTx).(*gorm.DB)
	if !ok {
		return nil
	}
	return tx.Rollback().Error
}
