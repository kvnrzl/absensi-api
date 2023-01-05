package mysql

import (
	"context"

	"absensi-api.com/domain/attendance/repository"
	"absensi-api.com/model"
	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct{}

func NewAttendanceRepositoryImpl() repository.AttendanceRepository {
	return &AttendanceRepositoryImpl{}
}

func (r *AttendanceRepositoryImpl) Create(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error) {
	if err := db.WithContext(ctx).Create(attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *AttendanceRepositoryImpl) FindByUserID(ctx context.Context, db *gorm.DB, userID uint) ([]*model.Attendance, error) {
	var attendances []*model.Attendance
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *AttendanceRepositoryImpl) FindLatestAttendanceInfoByUserID(ctx context.Context, db *gorm.DB, userID uint) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Last(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepositoryImpl) Update(ctx context.Context, db *gorm.DB, attendance *model.Attendance) (*model.Attendance, error) {
	if err := db.WithContext(ctx).Where("id = ?", attendance.ID).Updates(&attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}
