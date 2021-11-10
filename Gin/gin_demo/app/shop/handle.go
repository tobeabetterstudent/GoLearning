/*用来定义shop的router的处理函数*/
package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	name := c.DefaultQuery("name","书包")
	price  := c.DefaultQuery("price", "470")
	c.String(http.StatusOK, name + " is " + price + "元") 
}

func buyHandler(c *gin.Context) {
	buyer := c.DefaultQuery("buyer","李四")
	c.String(http.StatusOK, buyer) 
}
