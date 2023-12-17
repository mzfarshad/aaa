package jwt

import (
	"web-service-gin/config"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type TokenUser struct {
	Email    string
	UserType string
}

func NewAccessToken(email, userType string) (string, error) {
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, gojwt.MapClaims{
		"email":     email,
		"user_type": userType,
	})
	secretKey := []byte(config.Get().JWT().SecretKey)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
