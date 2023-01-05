package model

import "errors"

var (
	ErrInvalidJsonRequest          = errors.New("Invalid json request")
	ErrUsernameAlreadyExist        = errors.New("Username telah terdaftar")
	ErrUsernameOrPasswordWrong     = errors.New("Username atau password salah")
	ErrAttendanceAlreadyCheckedIn  = errors.New("User telah melakukan check in")
	ErrAttendanceAlreadyCheckedOut = errors.New("User telah melakukan check out")
	ErrAttendanceNotCheckedIn      = errors.New("User belum melakukan check in")
	ErrUserNotLoggedIn             = errors.New("User belum login")
	ErrTokenNotValid               = errors.New("Token tidak valid")
)
