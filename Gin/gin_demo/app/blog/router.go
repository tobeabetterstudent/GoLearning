/*用来定义blog的router相关信息*/
package blog


import (
	"github.com/gin-gonic/gin"
)

func SetRouters(e *gin.Engine) {
	e.GET("/person", personHandler)
	e.GET("/comment", commentHandler)
}