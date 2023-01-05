package repository

import (
	"context"

	"absensi-api.com/model"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	// Create
	Create(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error)

	// Find All Attendance
	FindAll(ctx context.Context, db *gorm.DB) ([]*model.Attendance, error)

	// Find Attendance By UserID
	FindByUserID(ctx context.Context, db *gorm.DB, userID uint) ([]*model.Attendance, error)

	// Find the latest attendance information by user id
	FindLatestAttendanceInfoByUserID(ctx context.Context, db *gorm.DB, userID uint) (*model.Attendance, error)

	// Update
	Update(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error)
}
