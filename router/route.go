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

	router.GET("/", func(r *gin.Context) {
		r.JSON(200, "Absensi API service is ready cok!")
	})

	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
		api.POST("/logout", userController.Logout)
		api.Use(authMiddleware())
		{
			attendances := api.Group("/attendances")
			{
				attendances.POST("", attendanceController.CheckIn)
				attendances.PATCH("", attendanceController.CheckOut)
				attendances.GET("", attendanceController.GetAllHistoriesAttendance)
			}
			activities := api.Group("/activities")
			{
				activities.GET("", activityController.GetActivities)
				activities.GET("/filter", activityController.GetActivitiesByDate)
				activities.GET("/all", activityController.GetAllActivities)
				activities.Use(attendanceMiddleware())
				{
					activities.POST("", activityController.CreateActivity)
					activities.PATCH("/:id", activityController.UpdateActivity)
					activities.DELETE("/:id", activityController.DeleteActivity)
				}
			}
		}
	}

	return router
}
