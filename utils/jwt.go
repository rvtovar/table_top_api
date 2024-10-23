package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	salt, err := OpenSalt()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(salt))
}
