package handler

import (
	"net/http"
	"strconv"
	"web-service-gin/api/middlewares"
	"web-service-gin/api/presenter"
	"web-service-gin/models"
	"web-service-gin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func FollowUser(ctx *gin.Context) {
	strID := ctx.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid param id"))
		return
	}
	token, exist := ctx.Get(middlewares.AuthenticatedUserKey)
	if !exist {
		ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("please sing up"))
		return
	}
	user := token.(*jwt.TokenUser)
	if err := models.FindById(uint(id)); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("not found user"))
		return
	}
	following := &models.Following{
		UserID:         user.Id,
		FollowedUserID: uint(id),
	}
	if err := following.CreateFollower(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess("success following"))
}

func Profile(ctx *gin.Context) {
	strId := ctx.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid id"))
		return
	}
	if err := models.FindById(uint(id)); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("not found user"))
		return
	}
	profile := new(models.Profile)
	userProfile, err := profile.FindProfile(uint(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()).
			AppendMessages("try again"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(userProfile))

}
