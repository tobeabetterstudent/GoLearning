/*用来定义shop的router相关信息*/
package shop

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(e *gin.Engine) {
	e.GET("/goods", goodsHandler)
	e.GET("/buy", buyHandler)
}