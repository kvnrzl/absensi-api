package controller

import (
	"absensi-api.com/domain/user/service"
	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := u.userService.Register(c, &user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (u *UserControllerImpl) Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := u.userService.Login(c, &user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// cookie := http.Cookie{
	// 	Name:     "token",
	// 	Value:    token,
	// 	MaxAge:   int(time.Hour * 24),
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Expires:  time.Now().Add(24 * time.Hour),
	// }

	// http.SetCookie(c.Writer, &cookie)

	c.JSON(200, token)
}

func (u *UserControllerImpl) Logout(c *gin.Context) {
	// name, err := c.Cookie("token")
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// cookie := http.Cookie{
	// 	Name:     name,
	// 	Value:    "",
	// 	MaxAge:   -1,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Expires:  time.Now().Add(-time.Hour),
	// }

	// http.SetCookie(c.Writer, &cookie)

	c.JSON(200, "Logout success")
}
