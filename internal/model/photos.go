package model

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	PhotoURL  string    `json:"url"`
	Caption   string    `json:"caption"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time	
	UpdatedAt time.Time
}
