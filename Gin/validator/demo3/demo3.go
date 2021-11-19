package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Info struct{
	// 这里tag:binding中的约束timing就是一个自定义的约束
	CreateTime time.Time `form:"create_time" binding:"required,timing" time_format:"2006-01-02"`
 	UpdateTime time.Time `form:"update_time" binding:"required,timing" time_format:"2006-01-02"`
}

// timing 为timging约束执行的函数
func timing(f1 validator.FieldLevel) bool {
	if date, ok := f1.Field().Interface().(time.Time); ok{
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// getTime GET("/time")的handle
func getTime(c *gin.Context) {
	var info Info
	// 数据模型绑定查询字符串验证
	var err error
	if err = c.ShouldBind(&info); err == nil {
	 	c.JSON(http.StatusOK, gin.H{"message": "time are valid!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func main() {
	e := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", timing)
		if err != nil {
			fmt.Println("Register Constraint Timing Succeed")
		}
	}
	e.GET("/time", getTime)
	e.Run(":8080")
}