/*
1.首先使用Include函数将所有app的SetRouters函数注册进来
2.再调用InitRouters函数 将所有的路由注册到engine
*/
package routers

import (
	"github.com/gin-gonic/gin"
)

// Option是SetRouters函数的原型 将之type为Option 使得Include函数接受不定参数
type Option func(*gin.Engine)

// options是Option的slice 初始化为空 用于装载所有的SetRouters函数
var options = []Option{}

func Include(opts ...Option) {
	// opts...的实际类型就是[]Option 因此直接append
	options = append(options,opts...)
}

func InitRouters() *gin.Engine{
	// 首先创建路由
	e := gin.Default()
	for _, opt := range options {
		opt(e)
	}
	return e
}