package router

import (
	"fmt"

	"absensi-api.com/config"
	"absensi-api.com/helper"
	"absensi-api.com/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtString, err := c.Cookie("token")
		if err != nil {
			helper.ResponseUnauthorized(c, model.ErrUserNotLoggedIn.Error())
			c.Abort()
			return
		}

		_, err = jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			// Verify that the token is signed with the correct secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.SECRET_KEY), nil
		})

		if err != nil {
			helper.ResponseUnauthorized(c, model.ErrTokenNotValid.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
