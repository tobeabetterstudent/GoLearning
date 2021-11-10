// routes group是为了管理一些相同的URL
package main

import (
   "github.com/gin-gonic/gin"
   "fmt"
)

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
 }
 
 func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
 }

func example1() {
	r := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
	   v1.GET("/login", login)
	   v1.GET("submit", submit)
	}
	// 路由组2 ，处理POST请求
	v2 := r.Group("/v2")
	{
	   v2.POST("/login", login)
	   v2.POST("/submit", submit)
	}
	r.Run(":8000")
}

func main() {
	example1()
}