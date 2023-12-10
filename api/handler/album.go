package handler

import (
	"net/http"
	"web-service-gin/api/presenter"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responsde with the list of all album as JSON.
func GetAlbums(c *gin.Context) {

	var list []models.Album
	title := c.Query("title")

	if title != "" {
		for _, album := range models.Albums {
			if album.Title == title {
				list = append(list, album)
			}
		}
		c.IndentedJSON(http.StatusOK, presenter.Response{Data: list, IsSuccess: true})
		return
	}

	c.IndentedJSON(http.StatusOK, presenter.Response{Data: models.Albums, IsSuccess: true})
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(c *gin.Context) {
	var newAlbum models.Album
	//call bindjson to bind the recived json to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, presenter.Response{
			IsSuccess: false,
			Messages:  []string{"invalid body"},
		})
		return
	}
	//add newAlbum to slice
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, presenter.Response{
		Data:      newAlbum,
		IsSuccess: true,
		Messages:  []string{"successfully created"},
	})
}

/*
	GetAlbumByID locates the album whose ID value matches the id

parameter sent by the client, then returns that album as a response.
*/
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	/* Loop over the list of models.Albums, looking for
	   an album whose ID value matches the parameter.*/
	for _, album := range models.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, presenter.Response{
				Data:      album,
				IsSuccess: true,
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, presenter.Response{
		IsSuccess: false,
		Messages:  []string{"album not found"},
	})
}
