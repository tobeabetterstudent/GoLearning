// @program: gin_demo2在v2路由组的handles
// @author: aslanwang
// @create: 2021-11-13

package	v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"v2" : "AddMember"})
} 

func AddProduct(c *gin.Context)  {
	// 获取 Get 参数
	name  := c.DefaultQuery("name","铅笔")
	price := c.DefaultQuery("price", "100")

	c.JSON(200, gin.H{
		"v2"    : "AddProduct",
		"name"  : name,
		"price" : price,
	})
}