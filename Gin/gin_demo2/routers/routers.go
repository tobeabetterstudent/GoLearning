// @program: gin_demo2的路由包 在这里向engine中添加路由信息
// @author: aslanwang
// @create: 2021-11-13

package	routers


import (
	"github.com/gin-gonic/gin"
	"GoLearning/Gin/gin_demo2/controller/v1"
	"GoLearning/Gin/gin_demo2/controller/v2"
	"GoLearning/Gin/gin_demo2/common"
	"GoLearning/Gin/gin_demo2/myvalidator"
	//"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"strconv"
	"net/url"
)

func Sign(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name"  : []string{c.Query("name")},
		"price" : []string{c.Query("price")},
		"ts"    : []string{ts},
	}
	res["sn"] = common.CreateSign(params)
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
	// 将自定义tag加入到validator设置中去
	config := &validator.Config{TagName: "validate"}
	v := validator.New(config)
	// 注册tag validate的校验名与方法
	v.RegisterValidation("NameValid", myvalidator.NameValid)
	// RegisterValidation将tag 与 对应的处理函数 映射到map[string]Func中 这样根据tag就可以直接调用对应的处理函数
	// 其中: type Func func(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool
}