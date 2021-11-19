package main

import (
	"github.com/gin-gonic/gin"
	"GoLearning/Gin/gin_demo2/routers"
	"GoLearning/Gin/gin_demo2/config"
	//"GoLearning/Gin/gin_demo2/middleware"
)

func main() {
	e := gin.Default()
	//e.Use(middleware.LoggerToFile())
	routers.InitRouters(e)
	e.Run(config.PORT)
}