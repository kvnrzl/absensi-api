package router

import (
	activityController "absensi-api.com/domain/activity/controller"
	attendanceController "absensi-api.com/domain/attendance/controller"
	userController "absensi-api.com/domain/user/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userController userController.UserController,
	activityController activityController.ActivityController,
	attendanceController attendanceController.AttendanceController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
		api.POST("/logout", userController.Logout)

		activities := api.Group("/activities")
		{
			activities.GET("", activityController.GetAllActivities)
			activities.POST("", activityController.CreateActivity)
			activities.GET("/:id", activityController.GetActivityByID)
			activities.PATCH("/:id", activityController.UpdateActivity)
			activities.DELETE("/:id", activityController.DeleteActivity)
		}

		attendances := api.Group("/attendances")
		{
			attendances.POST("", attendanceController.CheckIn)
			attendances.PATCH("", attendanceController.CheckOut)
			attendances.GET("", attendanceController.GetAllHistoriesAttendance)
		}
	}

	return router
}
