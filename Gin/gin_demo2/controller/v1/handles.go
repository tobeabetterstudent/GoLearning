// @program: gin_demo2在v1路由组的handles
// @author: aslanwang
// @create: 2021-11-13

package	v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context)  {
	// 获取 Get 参数
	name  := c.DefaultQuery("name","书包")
	price := c.DefaultQuery("price", "100")

	c.JSON(http.StatusOK, gin.H{
			"v1"    : "AddMember",
			"name"  : name,
			"price" : price,
	})
} 

func AddProduct(c *gin.Context)  {
	// 获取 Get 参数
	name :=  c.DefaultQuery("name","钢笔")
	price := c.DefaultQuery("price", "100")

	c.JSON(200, gin.H{
		"v1"    : "AddProduct",
		"name"  : name,
		"price" : price,
	})
}