package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 分组路由
	v1 := r.Group("v1")
	{
		v1.GET("/h", func(g *gin.Context) {
			g.JSON(200, gin.H{
				"h": 1,
			})
		})
		v1.GET("/e", func(g *gin.Context) {
			g.JSON(200, gin.H{
				"e": 2,
			})
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/l", func(g *gin.Context) {
			g.JSON(200, gin.H{
				"l": 3,
			})
		})
		v2.GET("o", func(g *gin.Context) {
			g.JSON(200, gin.H{
				"o": 4,
			})
		})
	}
	r.Run()
}
