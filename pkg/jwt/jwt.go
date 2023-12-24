package jwt

import (
	"fmt"
	"log"
	"web-service-gin/config"
	"web-service-gin/models"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type TokenUser struct {
	Email    string
	UserType string
}

func NewAccessToken(email string, userType models.UserType) (string, error) {
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

func (t TokenUser) SetTokenUser(authHeader string) (*TokenUser, error) {
	token, _, err := new(gojwt.Parser).ParseUnverified(authHeader, gojwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(gojwt.MapClaims); ok {
		t.Email = fmt.Sprint(claims["email"])
		t.UserType = fmt.Sprint(claims["user_type"])
		return &t, nil
	}
	return nil, err
}

func PrintTokenUser(t TokenUser) {
	log.Printf(" Email: %s , UserType : %s ", t.Email, t.UserType)
}
