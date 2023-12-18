package handler

import (
	"net/http"
	"web-service-gin/api/presenter"
	"web-service-gin/models"
	"web-service-gin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// SignIn
func SignIn(ctx *gin.Context) {
	var req presenter.LoginRequest
	//call bindjson to bind the recived json to newAlbum.
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid body").AppendMessage("test error"))
		return
	}
	// TODO: @Farshad
	// 0. Get user by email from database, or return error
	// 1. Check if req.Password == user.Password
	userType := "user"
	// 2. change token user type claim if user is admin
	// if user.IsAdmin {
	// 	userType = "admin"
	// }
	token, err := jwt.NewAccessToken(req.Email, userType)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	response := presenter.Token{Access: token}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response).AppendMessage("successfull login"))
}

// SignUp creates a user by email in case of no duplication
func SignUp(ctx *gin.Context) {
	// TODO: @Farshad implement me
	// 1. Get request body from user (email, password)
	// 2. Check if given email does not exist in database, or return error
	// 3. Create a new user in users table
	var req presenter.SignUpRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid body"))
		return
	}
	user := new(models.User)

	// if err := user.FindByEmail(req.Email); err != nil {
	// 	if !errors.Is(err, errors.New("email not found")) {
	// 		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
	// 		return
	// 	}
	// }
	// if user.ID > 0 { //user with request email exists
	// 	ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed("email already signed up, please login"))
	// 	return
	// }
	user.Email = req.Email
	user.Password = req.Password
	if err := user.Create(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	token, err := jwt.NewAccessToken(req.Email, "user")
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	response := presenter.Token{Access: token}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response).AppendMessage("successfully signed up"))
}
