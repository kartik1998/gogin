package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "AoK",
		})
	})
	server.Run(":8080")
}
