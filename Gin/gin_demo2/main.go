package main

import (
	"github.com/gin-gonic/gin"
	"GoLearning/Gin/gin_demo2/routers"
	"GoLearning/Gin/gin_demo2/config"
	"GoLearning/Gin/gin_demo2/common"
	"fmt"
)

func main() {
	fmt.Println(common.GetTimeUnix())
	e := gin.Default()
	routers.InitRouters(e)
	e.Run(config.PORT)
}