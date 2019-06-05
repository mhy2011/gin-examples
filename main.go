package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	// 启动服务,默认端口为8080
	r.Run()
}
