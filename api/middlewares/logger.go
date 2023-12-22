package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger(ctx *gin.Context) {
	log.Println("")
	log.Println("Middleware : Call API")
	log.Println("")
	ctx.Next()
}
