package model

import (
	"github.com/google/uuid"
	"time"
)

type SocialMedia struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name           string
	SocialMediaURL string
	UserID         uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
