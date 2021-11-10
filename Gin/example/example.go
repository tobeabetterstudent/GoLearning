package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
)

// demoByNet 直接使用net/http包来处理http请求
func demoByNet() {
	// 注册一个路由('/') 及这个路由的handler到DefaultServeMux中
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
	// ln, err := net.Listen("tcp", addr)做了初试化了socket, bind, listen的操作.
	// rw, e := l.Accept()进行accept, 等待客户端进行连接
	// go c.serve(ctx) 启动新的goroutine来处理本次请求. 同时主goroutine继续等待客户端连接, 进行高并发操作
	// h, _ := mux.Handler(r) 获取注册的路由, 然后拿到这个路由的handler, 然后将处理结果返回给客户端
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("start http server fail:", err)
    }
}

// example1 使用gin的第一个example
func example1() {
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

// example2 使用gin的第二个example
// 可以通过Context的Param方法来获取API参数
// 一般带API参数的URL形如-> localhost:8080/user/name/action
func example2() {
	r := gin.Default()
	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name + " is " + action)
	})
	 // 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

// example3 使用gin的第三个example
// 可以通过Context的DefaultQuery()或Query()方法来获取URL参数
// 一般带参数的URL形如-> localhost:8080/user?name=xxx&age=yyy
func example3() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name","zhangsan")
		c.String(http.StatusOK, name)
	})
	 // 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

// example4 使用gin的第四个example
// 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
// 表单传输为post请求 该参数在http请求的body中 不在URL中
func example4() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type","post")
		username := c.DefaultPostForm("username","zhangsan")
		password := c.DefaultPostForm("password","123456")
		c.String(http.StatusOK, username + " " +  password + " " + types)
	})
	 // 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

// example5 使用gin的第五个example
// 使用FormFile上传文件 并自增函数以限制文件的类型和大小
func example5() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
			return
		}
		if headers.Size > (2 << 20) {
			log.Printf("file size is too big: %v", headers.Size)
			return
		}
		if headers.Header.Get("Content-Type") != "image/jpeg"{
			log.Printf("file type is error! Need jpg")
			return
		}
		//c.SaveUploadedFile(headers, "./video/"+headers.Filename)
        c.String(http.StatusOK, "文件上传成功 : " + headers.Filename)
	})
	 // 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	example5()
}