package Server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 定义路由
var R = gin.Default()

func Gin_init() {
	// 允许所有跨域请求

	R.Use(cors.Default())
	R.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

}
