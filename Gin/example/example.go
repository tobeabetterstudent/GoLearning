package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)


// 直接使用net/http包来处理http请求
func demoByNet() {
	// 注册一个路由('/') 及这个路由的handler到DefaultServeMux中
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
	// ln, err := net.Listen("tcp", addr)做了初试化了socket, bind, listen的操作.
	// rw, e := l.Accept()进行accept, 等待客户端进行连接
	// go c.serve(ctx) 启动新的goroutine来处理本次请求. 同时主goroutine继续等待客户端连接, 进行高并发操作
	// h, _ := mux.Handler(r) 获取注册的路由, 然后拿到这个路由的handler, 然后将处理结果返回给客户端
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("start http server fail:", err)
    }
}


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