package repository

import (
	"context"

	"absensi-api.com/model"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error)
	FindByUserID(ctx context.Context, db *gorm.DB, userID uint) ([]*model.Attendance, error)
	FindLatestAttendanceInfoByUserID(ctx context.Context, db *gorm.DB, userID uint) (*model.Attendance, error)
	Update(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error)
}
