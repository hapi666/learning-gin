package main

import (
	"github.com/gin-gonic/gin"
)

// 起一个服务
func main() {
	g := gin.Default()
	g.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "demo1",
		})
	})
	g.Run()
}
