package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID
	PhotoID   uuid.UUID
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
