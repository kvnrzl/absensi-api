package controller

import "github.com/gin-gonic/gin"

type AttendanceController interface {
	CheckIn(c *gin.Context)
	CheckOut(c *gin.Context)
	GetAllHistoriesAttendance(c *gin.Context)
}
