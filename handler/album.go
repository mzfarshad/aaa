package handler

import (
	"net/http"
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
		c.IndentedJSON(http.StatusOK, list)
		return
	}

	c.IndentedJSON(http.StatusOK, models.Albums)
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(c *gin.Context) {
	var newAlbum models.Album
	//call bindjson to bind the recived json to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
		return
	}
	//add newAlbum to slice
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
	GetAlbumByID locates the album whose ID value matches the id

parameter sent by the client, then returns that album as a response.
*/
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	/* Loop over the list of models.Albums, looking for
	   an album whose ID value matches the parameter.*/
	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
