package service

import (
	"context"

	"absensi-api.com/model"
)

type AttendanceService interface {
	// Check In
	CheckIn(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error)

	// Check Out
	CheckOut(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error)

	// Get Histories of Attendance
	GetAllHistoriesAttendance(ctx context.Context, userID uint) ([]*model.Attendance, error)
}
