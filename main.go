package main

import (
	"web-service-gin/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.CreateNewAlbum)

	router.Run("localhost:8080")
}
