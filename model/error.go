package model

import "errors"

var (
	ErrUsernameAlreadyExist        = errors.New("username already exist")
	ErrAttendanceAlreadyCheckedIn  = errors.New("you already checked in")
	ErrAttendanceAlreadyCheckedOut = errors.New("you already checked out")
	ErrAttendanceNotCheckedIn      = errors.New("you haven't checked in yet")
)
