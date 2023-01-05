package controller

import (
	"errors"
	"net/http"
	"time"

	"absensi-api.com/domain/user/service"
	"absensi-api.com/helper"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	_, err := u.userService.Register(c, &user)
	if err != nil {
		if errors.Is(err, model.ErrInvalidJsonRequest) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrUsernameAlreadyExist) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	logrus.Infof("User created: %v", user)
	helper.ResponseCreated(c, "Register success")
}

func (u *UserControllerImpl) Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	token, err := u.userService.Login(c, &user)
	if err != nil {
		if errors.Is(err, model.ErrInvalidJsonRequest) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		if errors.Is(err, model.ErrUsernameOrPasswordWrong) {
			logrus.Error(err)
			helper.ResponseBadRequest(c, err.Error())
			return
		}

		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(c.Writer, &cookie)

	logrus.Infof("User logged in: %v", user)
	helper.ResponseOK(c, "Login success")
}

func (u *UserControllerImpl) Logout(c *gin.Context) {
	attCookie := http.Cookie{
		Name:     "attendance",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Second),
	}
	http.SetCookie(c.Writer, &attCookie)

	tokenCookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour),
	}
	http.SetCookie(c.Writer, &tokenCookie)

	logrus.Infof("User logged out")
	helper.ResponseOK(c, "Logout success")
}
