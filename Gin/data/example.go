// @program: gin数据解析和绑定
// @author: aslanwang
// @create: 2021-11-11
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
) 

// 定义接受数据的结构体
type Login struct{
	// binding:"required" 表明该数据是必须字段 若解析为空值则报错
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// handleJSON 处理客户端POST请求中req body数据为json格式 
// postman中Body输入raw {"username":"root","password":"root"} 格式选择JSON
// POST发送json数据时 header中的Content-Type:application/josn
// POST提交表单数据时 header中的Content-Type:application/x-www-form-urlencoded
// 这两者在浏览器上的表现是不一致的
func handleJSON() {
	e := gin.Default()
	e.POST("loginJSON",func(c *gin.Context){
		var input Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		// body中形如{"username":"root","password":"root"}
		if err := c.ShouldBindJSON(&input); err != nil {
			// 封装错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if input.Username != "root" || input.Password != "root"{
			// 封装登录信息错误
			c.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status": 200})
	})
	e.Run(":8080")
}

// handleJSON 处理客户端POST请求提交的表单数据
// 
func handleFormData() {
	e := gin.Default()
	e.POST("loginForm",func(c *gin.Context){
		var input Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		// body中形如{"username":"root","password":"root"}
		if err := c.ShouldBind(&input); err != nil {
			// 封装错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if input.Username != "root" || input.Password != "root"{
			// 封装登录信息错误
			c.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status": 200})
	})
	e.Run(":8080")
}

func main(){
	//handleJSON()
	handleFormData()
}


 