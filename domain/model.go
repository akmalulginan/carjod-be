package domain

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        string          `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
