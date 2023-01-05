package controller

import (
	"errors"
	"strconv"

	"absensi-api.com/domain/activity/service"
	"absensi-api.com/helper"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
)

type ActivityControllerImpl struct {
	activityService service.ActivityService
}

func NewActivityControllerImpl(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		activityService: activityService,
	}
}

func (a *ActivityControllerImpl) CreateActivity(c *gin.Context) {
	var activity model.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	activity.UserID = uint(userID)

	res, err := a.activityService.Save(c, &activity)
	if err != nil {
		if errors.Is(err, model.ErrInvalidJsonRequest) {
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseCreated(c, res)
}

func (a *ActivityControllerImpl) UpdateActivity(c *gin.Context) {
	var activity model.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}
	activity.ID = uint(id)

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)
	activity.UserID = uint(userID)

	res, err := a.activityService.Update(c, &activity)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, res)
}

func (a *ActivityControllerImpl) DeleteActivity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)

	if err := a.activityService.Delete(c, uint(id), uint(userID)); err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, "Activity deleted")
}

func (a *ActivityControllerImpl) GetAllActivities(c *gin.Context) {
	res, err := a.activityService.FindAll(c)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, res)
}

func (a *ActivityControllerImpl) GetActivities(c *gin.Context) {
	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)

	res, err := a.activityService.FindByUserID(c, uint(userID))
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, res)
}

func (a *ActivityControllerImpl) GetActivitiesByDate(c *gin.Context) {
	date := c.Query("startdate")

	jwtString, _ := c.Cookie("token")
	userID, _, _, _ := helper.ExtractCookie(jwtString)

	res, err := a.activityService.FindByDate(c, uint(userID), date)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseOK(c, res)
}
