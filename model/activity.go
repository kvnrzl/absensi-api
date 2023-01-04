package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"user_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartTime   string         `json:"start_time"`
	EndTime     string         `json:"end_time"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
}
