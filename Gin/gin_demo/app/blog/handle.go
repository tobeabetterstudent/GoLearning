/*用来定义blog的router的处理函数*/
package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func personHandler(c *gin.Context) {
	name := c.DefaultQuery("name","张三")
	age  := c.DefaultQuery("age", "18")
	c.String(http.StatusOK, name + " is " + age + "岁") 
}

func commentHandler(c *gin.Context) {
	text := c.DefaultQuery("text","今天天气很好")
	c.String(http.StatusOK, text) 
}
