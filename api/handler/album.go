package handler

import (
	"net/http"
	"web-service-gin/api/presenter"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responsde with the list of all album as JSON.
func GetAlbums(ctx *gin.Context) {

	var list []models.Album
	title := ctx.Query("title")

	if title != "" {
		for _, album := range models.Albums {
			if album.Title == title {
				list = append(list, album)
			}
		}
		ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(list))
		return
	}

	ctx.IndentedJSON(http.StatusOK, presenter.NewSuccess(models.Albums))
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(ctx *gin.Context) {
	var newAlbum models.Album
	//call bindjson to bind the recived json to newAlbum.
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, presenter.
			NewFailed("invalid body").
			AppendMessage("test error"))
		return
	}
	//add newAlbum to slice
	models.Albums = append(models.Albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, presenter.
		NewSuccess(newAlbum).
		AppendMessage("successfully created").
		AppendMessage("test message"),
	)
}

/*
	GetAlbumByID locates the album whose ID value matches the id

parameter sent by the client, then returns that album as a response.
*/
func GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	/* Loop over the list of models.Albums, looking for
	   an album whose ID value matches the parameter.*/
	for _, album := range models.Albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, presenter.Response{
				Data:      album,
				IsSuccess: true,
			})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, presenter.Response{
		IsSuccess: false,
		Messages:  []string{"album not found"},
	})
}
