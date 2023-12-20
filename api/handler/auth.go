package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"web-service-gin/api/presenter"
	"web-service-gin/models"
	"web-service-gin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// SignIn
func SignIn(ctx *gin.Context) {
	var req presenter.SignInRequest
	//call bindjson to bind the recived json to newAlbum.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		messages := strings.Split(err.Error(), "\n")
		res := presenter.NewFailed("failed to bind json").AppendMessages(messages...)
		ctx.IndentedJSON(http.StatusBadRequest, res)
		return
	}

	user := new(models.User)

	if err := user.FindByEmail(req.Email); err != nil {
		if !errors.Is(err, models.ErrEmailNotFound) {
			ctx.IndentedJSON(http.StatusInternalServerError, presenter.
				NewFailed("something is wrong, please try again"))
			log.Print(err)
			return
		}
		ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("email not found").
			AppendMessages("please try again"))
		return
	}

	// if req.Password != user.Password {
	// 	ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("invalid password").
	// 		AppendMessages("please try again"))
	// 	return
	// }
	if !models.CheckPasswordHash(req.Password, user.Password) {
		ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("invalid password").
			AppendMessages("please try again"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess("successfully signed in"))

	userType := models.UserTypeUser
	// 2. change token user type claim if user is admin
	// if user.IsAdmin {
	// 	userType = models.UserTypeAdmin
	// }
	token, err := jwt.NewAccessToken(req.Email, userType)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	response := presenter.Token{Access: token}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response).AppendMessages("successfully signed in"))
}

// SignUp creates a user by email in case of no duplication
func SignUp(ctx *gin.Context) {
	var req presenter.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid body"))
		return
	}

	user := new(models.User)
	if err := user.FindByEmail(req.Email); err != nil {
		if !errors.Is(err, models.ErrEmailNotFound) {
			ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
			return
		}
	}
	if user.ID > 0 { //user with request email exists
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed("email already signed up, please sign in"))
		return
	}
	user.Email = req.Email
	user.Password = req.Password
	if err := user.Create(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	token, err := jwt.NewAccessToken(req.Email, models.UserTypeUser)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	response := presenter.Token{Access: token}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response).AppendMessages("successfully signed up"))
}
