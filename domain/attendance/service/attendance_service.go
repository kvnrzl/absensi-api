package service

import (
	"context"

	"absensi-api.com/model"
)

type AttendanceService interface {
	CheckIn(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error)
	CheckOut(ctx context.Context, attendance *model.Attendance) (*model.Attendance, error)
	GetAllHistoriesAttendance(ctx context.Context, userID uint) ([]*model.Attendance, error)
}
