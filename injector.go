//go:build wireinject
// +build wireinject

package main

import (
	"absensi-api.com/database"
	activityController "absensi-api.com/domain/activity/controller"
	activityRepo "absensi-api.com/domain/activity/repository/mysql"
	activityService "absensi-api.com/domain/activity/service"
	attendanceController "absensi-api.com/domain/attendance/controller"
	attendanceRepo "absensi-api.com/domain/attendance/repository/mysql"
	attendanceService "absensi-api.com/domain/attendance/service"
	userController "absensi-api.com/domain/user/controller"
	userRepo "absensi-api.com/domain/user/repository/mysql"
	userService "absensi-api.com/domain/user/service"
	"absensi-api.com/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		database.InitDBMysql,
		validator.New,

		userRepo.NewUserRepositoryImpl,
		userService.NewUserServiceImpl,
		userController.NewUserControllerImpl,

		activityRepo.NewActivityRepositoryImpl,
		activityService.NewActivityServiceImpl,
		activityController.NewActivityControllerImpl,

		attendanceRepo.NewAttendanceRepositoryImpl,
		attendanceService.NewAttendanceServiceImpl,
		attendanceController.NewAttendanceControllerImpl,

		router.SetupRouter,
	)

	return nil
}
