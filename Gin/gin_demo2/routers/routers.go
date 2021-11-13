// @program: gin_demo2的路由包 在这里向engine中添加路由信息
// @author: aslanwang
// @create: 2021-11-13

package	routers


import (
	"github.com/gin-gonic/gin"
	"GoLearning/Gin/gin_demo2/controller/v1"
	"GoLearning/Gin/gin_demo2/controller/v2"
	"GoLearning/Gin/gin_demo2/common"
	"strconv"
	"net/url"
	"fmt"
)

func Sign(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name"  : []string{"a"},
		"price" : []string{"10"},
		"ts"    : []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	fmt.Println(res["sn"])
	res["ts"] = ts
	common.MakeJSON("200", "", res, c)
}

func InitRouters(e *gin.Engine) {
	// 在这个路由下根据url生成签名
	e.GET("/sn", Sign)
	// 生成签名后一般来说要把这个签名留给前端 再下次请求的时候再根据url生成签名 并和url中的签名参数进行校验
	group1 := e.Group("/v1")
	{
		group1.Any("/product/add", v1.AddProduct)
		group1.Any("/member/add", v1.AddMember)
	}
	group2 := e.Group("/v2", common.VerifySign)
	{
		group2.Any("/product/add", v2.AddProduct)
		group2.Any("/member/add", v2.AddMember)
	}
}