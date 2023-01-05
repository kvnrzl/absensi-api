package controller

import "github.com/gin-gonic/gin"

type ActivityController interface {
	CreateActivity(c *gin.Context)
	UpdateActivity(c *gin.Context)
	DeleteActivity(c *gin.Context)
	GetAllActivities(c *gin.Context)
	GetActivities(c *gin.Context)
	GetActivitiesByDate(c *gin.Context)
}
