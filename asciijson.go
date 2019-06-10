package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/json", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})
	engine.Run()
}
