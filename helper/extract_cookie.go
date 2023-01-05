package helper

import (
	"fmt"

	"absensi-api.com/config"
	"github.com/dgrijalva/jwt-go"
)

func ExtractCookie(jwtString string) (float64, string, string, string) {
	token, _ := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the token is signed with the correct secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})

	id := token.Claims.(jwt.MapClaims)["id"].(float64)
	name := token.Claims.(jwt.MapClaims)["name"].(string)
	username := token.Claims.(jwt.MapClaims)["username"].(string)
	role := token.Claims.(jwt.MapClaims)["role"].(string)

	return id, name, username, role
}
