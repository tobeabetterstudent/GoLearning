package main

import (
	"GoLearning/Gin/gin_demo/routers"
	"GoLearning/Gin/gin_demo/app/shop"
	"GoLearning/Gin/gin_demo/app/blog"
	"fmt"
)
func main() {
	// 加载多个APP的路由配置
	routers.Include(blog.SetRouters, shop.SetRouters)
	// 初始化路由
    e := routers.InitRouters()
	if err := e.Run(); err != nil {
        fmt.Println("startup service failed, err:%v\n", err)
    }
}