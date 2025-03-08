package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkeyfor/signing-tokens"

func GenerateToken(phone, password string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":    phone,
		"id":       id,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token.!")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, errors.New("invalid token.!")
	}

	validToken := parsedToken.Valid

	if !validToken {
		return -1, errors.New("invalid token.!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid token.!")
	}

	userId := int64(claims["id"].(float64))

	return userId, nil

}
