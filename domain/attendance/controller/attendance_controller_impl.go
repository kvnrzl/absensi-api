package controller

import (
	"errors"
	"net/http"
	"time"

	"absensi-api.com/domain/attendance/service"
	"absensi-api.com/helper"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrAttendanceAlreadyCheckedIn) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "attendance",
		Value:    "checked-in",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	logrus.Infof("Attendance checked in: %v", res)
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
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrAttendanceAlreadyCheckedOut) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		logrus.Error(err)
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

	logrus.Infof("Attendance checked out: %v", res)
	helper.ResponseOK(c, res)
}

func (a *AttendanceControllerImpl) GetAllHistoriesAttendance(c *gin.Context) {
	var attendance model.Attendance

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	attendance.UserID = uint(userID)

	res, err := a.attendanceService.GetAllHistoriesAttendance(c, attendance.UserID)
	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	logrus.Infof("Attendance histories: %v", res)
	helper.ResponseOK(c, res)
}
