package controller

import (
	"errors"
	"net/http"
	"time"

	"absensi-api.com/domain/attendance/service"
	"absensi-api.com/helper"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
)

type AttendanceControllerImpl struct {
	attendanceService service.AttendanceService
}

func NewAttendanceControllerImpl(attendanceService service.AttendanceService) AttendanceController {
	return &AttendanceControllerImpl{
		attendanceService: attendanceService,
	}
}

func (a *AttendanceControllerImpl) CheckIn(c *gin.Context) {
	var attendance model.Attendance

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	attendance.UserID = uint(userID)

	res, err := a.attendanceService.CheckIn(c, &attendance)
	if err != nil {
		if errors.Is(err, model.ErrInvalidJsonRequest) {
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrAttendanceAlreadyCheckedIn) {
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "attendance",
		Value:    "checked-in",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	helper.ResponseOK(c, res)
}

func (a *AttendanceControllerImpl) CheckOut(c *gin.Context) {
	var attendance model.Attendance

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	attendance.UserID = uint(userID)

	res, err := a.attendanceService.CheckOut(c, &attendance)
	if err != nil {
		if errors.Is(err, model.ErrInvalidJsonRequest) {
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrAttendanceAlreadyCheckedOut) {
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "attendance",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Second),
	}

	http.SetCookie(c.Writer, &cookie)

	helper.ResponseOK(c, res)
}

func (a *AttendanceControllerImpl) GetAllHistoriesAttendance(c *gin.Context) {
	var attendance model.Attendance

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	attendance.UserID = uint(userID)

	res, err := a.attendanceService.GetAllHistoriesAttendance(c, attendance.UserID)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, res)
}
