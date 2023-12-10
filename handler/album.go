package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "faryad", Artist: "dariush", Price: 23.4},
	{ID: "2", Title: "pariche", Artist: "moien", Price: 25.7},
	{ID: "3", Title: "talab", Artist: "ebi", Price: 20.5},
}

// GetAlbums responsde with the list of all album as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// CreateNewAlbum adds an album from json recived in the requst body.
func CreateNewAlbum(c *gin.Context) {
	var newAlbum album
	//call bindjson to bind the recived json to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
		return
	}
	//add newAlbum to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
	GetAlbumByID locates the album whose ID value matches the id

parameter sent by the client, then returns that album as a response.
*/
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	/* Loop over the list of albums, looking for
	   an album whose ID value matches the parameter.*/
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
