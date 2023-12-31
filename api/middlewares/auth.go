package middlewares

import (
	"log"
	"net/http"
	"strings"
	"web-service-gin/api/presenter"
	"web-service-gin/models"
	"web-service-gin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	AuthenticatedUserKey string = "Authenticated-AAA-User"
)

func Authenticate(ctx *gin.Context) {
	// TODO: @Farshad
	// Get the token of the user from the context header("Authorization")
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.Next()
		return
	}
	// 2. Remove the word "Bearer" from the authHeader variable, and get the pure token
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	//log.Println(authHeader)
	// 3. Search in our jwt package github and implement a method in pkg/jwt to convert a token to the TokenUser struct
	// 4. Set the context key ("Authenticated-AAA-User") to the result struct
	tokenUser, err := jwt.Validate(authHeader)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, presenter.NewFailed("invalid token"))
		ctx.Next()
		return
	}
	log.Println(tokenUser)
	ctx.Set(AuthenticatedUserKey, tokenUser)
	ctx.Next()
}

// OnlyUser should be called after authenticate.
func OnlyUser(ctx *gin.Context) {
	//TODO: @Farshad
	value, exist := ctx.Get(AuthenticatedUserKey)
	if !exist {
		ctx.IndentedJSON(http.StatusUnauthorized, presenter.NewFailed("required bearer token"))
		ctx.Next()
		return
	}
	// 1. Cast the above "value" variable to jwt.TokenUser struct (e.g, castedTokenUser variable).
	// 2. Check if castedTokenUser.UserType == models.UserTypeUser
	castedTokenUser := value.(*jwt.TokenUser)
	if castedTokenUser.UserType != string(models.UserTypeUser) {
		ctx.IndentedJSON(http.StatusUnauthorized, presenter.NewFailed("login is only allowed for the user"))
		ctx.Next()
		return
	}
	log.Println("Successful user login")
	ctx.Next()
}
