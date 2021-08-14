package main

import (
	"gogin-poc/controller"
	"gogin-poc/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "AoK",
		})
	})
	server.Run(":8080")
}
