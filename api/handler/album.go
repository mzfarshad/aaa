package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"web-service-gin/api/presenter"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responsde with the list of all album as JSON.
func GetAlbums(ctx *gin.Context) {
	var filter models.AlbumFilter
	if err := ctx.BindQuery(&filter); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			presenter.NewFailed("invalid query param").
				AppendMessages(err.Error()))
		return
	}

	var albums models.AlbumList
	if err := albums.Search(filter); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed(err.Error()))
		return
	}
	response := make(presenter.AlbumList, len(albums)).From(albums)
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response))
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(ctx *gin.Context) {
	var req presenter.CreateAlbumRequest
	//call bindjson to bind the recived json to newAlbum.
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.NewFailed("invalid body"))
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
	response := new(presenter.Album).From(album)
	ctx.IndentedJSON(http.StatusCreated, presenter.NewSuccess(response).AppendMessages("successfully created"))
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
		if errors.Is(err, models.ErrAlbumNotFound) {
			ctx.IndentedJSON(http.StatusNotFound, presenter.NewFailed("album not found"))
			return
		}
		log.Printf("failed to retrieve album >>> err: %v", err)
		ctx.IndentedJSON(http.StatusInternalServerError, presenter.NewFailed("please try again"))
		return
	}
	response := new(presenter.Album).From(album)
	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(response))
}
