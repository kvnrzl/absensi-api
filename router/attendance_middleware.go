package router

import (
	"absensi-api.com/helper"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
)

func attendanceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := c.Cookie("attendance")
		if err != nil {
			helper.ResponseBadRequest(c, model.ErrAttendanceNotCheckedIn.Error())
			c.Abort()
			return
		}

		if status != "checked-in" {
			helper.ResponseUnauthorized(c, model.ErrAttendanceNotCheckedIn.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
