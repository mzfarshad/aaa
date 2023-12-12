package handler

import (
	"net/http"
	"strconv"
	"web-service-gin/api/presenter"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responsde with the list of all album as JSON.
func GetAlbums(ctx *gin.Context) {
	title := ctx.Query("title")
	artist := ctx.Query("artist")

	var list models.AlbumList
	if err := list.Search(title, artist); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(list))
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(ctx *gin.Context) {
	var req presenter.CreateAlbumRequest
	//call bindjson to bind the recived json to newAlbum.
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid body").AppendMessage("test error"))
		return
	}
	album := &models.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}
	if err := album.Create(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	ctx.IndentedJSON(http.StatusCreated, presenter.NewSuccess(album).AppendMessage("successfully created"))
}

/*
GetAlbumByID locates the album whose ID value matches the id
parameter sent by the client, then returns that album as a response.
*/
func GetAlbumByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid id"))
		return
	}
	album := new(models.Album)
	if err := album.Find(uint(id)); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("album not found"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(album))
}
