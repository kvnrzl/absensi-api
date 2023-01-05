package model

import (
	"time"
)

type Attendance struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	CheckInTime  time.Time `json:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time"`
	// CreatedAt    time.Time      `json:"created_at"`
	// UpdatedAt    time.Time      `json:"updated_at"`
	// DeleteAt     gorm.DeletedAt `json:"delete_at"`
}
