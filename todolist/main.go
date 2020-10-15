package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("template/*")
	router.LoadHTMLFiles("template/template.html")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
