package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte("SECRET_KEY"))

	return t, err
}
