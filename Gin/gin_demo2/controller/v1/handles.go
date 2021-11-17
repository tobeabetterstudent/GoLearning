// @program: gin_demo2在v1路由组的handles
// @author: aslanwang
// @create: 2021-11-13

package	v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"GoLearning/Gin/gin_demo2/entity"
	"GoLearning/Gin/gin_demo2/alarm"
)

func AddMember(c *gin.Context)  {
	// 进行数据解析与校验
	member := new(entity.Member)
	// 这里通过ShouldBind自行查找合适的解析方式 如果是GET方法 则会去解析url;如果是POST则解析body 而body可以是form格式也可以是json格式...
	// GET方法一般是没有Content-Type这个请求头 这个请求头一般是POST数据携带在ReqBody中才会有 
	// 不论是GET传递url解析 还是POST传递的body解析 struct字段后的tag都会发挥作用
	if err := c.ShouldBind(member); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"v1"  : "AddMember",
			"msg" : "error!",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"v1"  : "AddMember",
		"name": member.Name,
		"age" : member.Age,
	})
} 

// hello 被AddProduct调用的内部方法 用于检查name是否为空 如果为空则抛出一个错误
func checkProductNameIsEmpty(name string) (err error) {
	if name == "" {
		err = alarm.WeChat("name 不能为空")
	}
	return
}

func AddProduct(c *gin.Context)  {
	// 获取 Get 参数
	name :=  c.Query("name")
	price := c.DefaultQuery("price", "100")
	if err := checkProductNameIsEmpty(name); err != nil{
		c.JSON(http.StatusForbidden, gin.H{
			"v1"  : "AddProduct",
			"msg" : err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"v1"    : "AddProduct",
		"name"  : name,
		"price" : price,
	})
}