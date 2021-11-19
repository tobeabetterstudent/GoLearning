// @program: 学习validator field级约束
// @author: aslanwang
// @time: 2021-11-19

package main

import (
	"fmt"
	
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"min=6,max=10"`
 	Age      uint8  `validate:"gte=1,lte=10"`
 	Sex      string `validate:"oneof=female male"`
}

func main() {
	// 创建一个验证器 下面再通过其Struct方法验证对象的字段是否符合tag约束
	v := validator.New()

	u1 := User{Username: "aslan", Age: 11, Sex: "null"}
	if err := v.Struct(u1); err != nil{
		fmt.Println(err)
	}
	
	u2 := User{Username: "aslanwang", Age: 8, Sex: "male"}
	if err := v.Struct(u2); err != nil{
		fmt.Println(err)
	}
}
