package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitsune2411/ginCRUD/config"
)

func main() {
	app := gin.Default()
	config.LoadEnv()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hellow hellow",
		})
	})
	app.Run()
}
