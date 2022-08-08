package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwtv4 "github.com/golang-jwt/jwt"
)

func CreateToken(userId int) (string, error) {
	claims := jwtv4.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expired after 1 hour
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	secret := os.Getenv("DB_SECRET")
	return token.SignedString([]byte(secret))
}

func ExtractToken(c gin.Context) (int, error) {
	user := c.MustGet("user").(*jwtv4.Token)
	if user.Valid {
		claims := user.Claims.(jwtv4.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId), nil
	}
	return 0, fmt.Errorf("token invalid")
}
