package main

import (
	"log"
	"web-service-gin/api/handler"
	"web-service-gin/api/middlewares"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %s", err)
	}
}

func main() {
	if err := models.ConnectToPostgres(); err != nil {
		panic("failed to connect database")
	}
	router := gin.Default()
	router.Use(middlewares.Authenticate)

	auth := router.Group("/auth")

	auth.POST("/signin", handler.SignIn)
	auth.POST("/signup", handler.SignUp)

	router.GET("/albums", middlewares.OnlyUser, handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.CreateNewAlbum)

	router.POST("/user/follows/:id", middlewares.OnlyUser, handler.FollowUser)
	router.GET("/user/follows", middlewares.OnlyUser, handler.Profile)

	router.Run("localhost:8080")
}
