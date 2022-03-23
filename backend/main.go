package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "post访问成功",
		})
	})

	r.GET("/toget", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "get访问成功",
		})
	})

	r.Run()
}
