package service

import (
	"context"
	"time"

	"absensi-api.com/domain/attendance/repository"
	"absensi-api.com/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AttendanceServiceImpl struct {
	attendanceRepository repository.AttendanceRepository
	db                   *gorm.DB
	validate             *validator.Validate
}

func NewAttendanceServiceImpl(attendanceRepository repository.AttendanceRepository, db *gorm.DB, validate *validator.Validate) AttendanceService {
	return &AttendanceServiceImpl{
		attendanceRepository: attendanceRepository,
		db:                   db,
		validate:             validate,
	}
}

// Check In
func (s *AttendanceServiceImpl) CheckIn(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error) {
	// validasi structnya
	if err := s.validate.Struct(attendance); err != nil {
		return nil, err
	}

	// cek apakah itu sudah pernah check in atau belum
	exist, err := s.attendanceRepository.FindLatestAttendanceInfoByUserID(ctx, s.db, attendance.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	attendance.CheckInTime = time.Now()

	if err == gorm.ErrRecordNotFound {
		return s.attendanceRepository.Create(ctx, s.db, attendance)
	}

	if exist.CheckInTime.Day() == time.Now().Day() {
		return nil, model.ErrAttendanceAlreadyCheckedIn
	}

	return s.attendanceRepository.Create(ctx, s.db, attendance)
}

// Check Out
func (s *AttendanceServiceImpl) CheckOut(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error) {
	// cek structnya
	if err := s.validate.Struct(attendance); err != nil {
		return nil, err
	}

	// cek apakah itu sudah pernah check out atau belum
	exist, err := s.attendanceRepository.FindLatestAttendanceInfoByUserID(ctx, s.db, attendance.UserID)
	if err != nil {
		return nil, err
	}

	if exist.CheckOutTime.Day() == time.Now().Day() {
		return nil, model.ErrAttendanceAlreadyCheckedOut
	}

	exist.CheckOutTime = time.Now()

	return s.attendanceRepository.Update(ctx, s.db, exist)
}

// Get Histories of Attendance
func (s *AttendanceServiceImpl) GetAllHistoriesAttendance(ctx context.Context, userID uint) ([]*model.Attendance, error) {
	return s.attendanceRepository.FindByUserID(ctx, s.db, userID)
}
