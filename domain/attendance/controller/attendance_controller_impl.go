package controller

import (
	"net/http"

	"absensi-api.com/domain/attendance/service"
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
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.attendanceService.CheckIn(c, &attendance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (a *AttendanceControllerImpl) CheckOut(c *gin.Context) {
	var attendance model.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.attendanceService.CheckOut(c, &attendance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (a *AttendanceControllerImpl) GetAllHistoriesAttendance(c *gin.Context) {
	var attendance model.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.attendanceService.GetAllHistoriesAttendance(c, attendance.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
