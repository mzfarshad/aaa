package jwt

import (
	"fmt"
	"web-service-gin/config"
	"web-service-gin/models"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type TokenUser struct {
	Email    string
	UserType string
	Id       uint
}

func NewAccessToken(email string, userType models.UserType, id uint) (string, error) {
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, gojwt.MapClaims{
		"email":     email,
		"user_type": userType,
		"id":        id,
	})
	secretKey := []byte(config.Get().JWT().SecretKey)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Validate(tkn string) (*TokenUser, error) {
	tokenUser := new(TokenUser)
	token, err := gojwt.Parse(tkn, func(token *gojwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Get().JWT().SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
		tokenUser.Email = (claim["email"]).(string)
		tokenUser.UserType = (claim["user_type"]).(string)
		tokenId := (claim["id"]).(float64)
		tokenUser.Id = uint(tokenId)
		return tokenUser, nil
	}
	return nil, fmt.Errorf("invalid token : %s", err)
}

func (t TokenUser) String() string {
	return fmt.Sprintf("Email : %s , UserType : %s , ID : %v", t.Email, t.UserType, t.Id)
}
