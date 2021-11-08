package main

import "github.com/gin-gonic/gin"


func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	r.GET("/ping", func(c *gin.Context) {
		// H即gin中对map[string]interface{}的重命名 便于使用
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}