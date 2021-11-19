[TOC]



# 请求参数校验——Validator包

## 1

### 目的：测试官方用例，首先掌握基本用法。

### 测试

- 结构体声明与代码见`demo1.go`，使用`Postman`模拟请求：

1. `post : http://127.0.0.1:8080/register`

   ​	报错：对结构体的所有字段都提示`"Key:` `'RegisterRequest.xxx' Error:Field validation for 'xxx' failed on the 'required' tag`

2. `post` 请求如下：

![image-20211119210730757](C:\Users\wangwei\AppData\Roaming\Typora\typora-user-images\image-20211119210730757.png)

​			报错：`email`字段不是一个合法邮箱，且`age`字段超出了最大限制。

## 2

### 目的：一个简单的示例：

```go
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

```

### 结果

​	`u1`报错如下：

```go
Key: 'User.Username' Error:Field validation for 'Username' failed on the 'min' tag
Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag     
Key: 'User.Sex' Error:Field validation for 'Sex' failed on the 'oneof' tag
```

## 3

### 目的：使用跨字段约束

### 简介

​	`validator`允许定义跨字段的约束，即约束该字段与其他字段之间的关系。这种约束实际上分为两种：

1.  这些参数字段都是同一个结构中的平级字段
2. 这些参数字段为结构中其他字段内部的字段，即约束的字段级次不一致。

约束语法很简单，只需要稍微修改一下。例如**相等约束**（`eq`）：

1. 如果是约束同级字段，则在后面添加一个`field`，使用`eqfield`定义字段间的相等约束。
2. 如果是更深层次的字段，在`field`之前还需要加上`cs`（可以理解为`cross-struct`），`eq`就变为`eqcsfield`。它们的参数值都是需要比较的字段名，内层的还需要加上字段的类型。

### 测试

​	测试代码如下：

```go
// @program: 跨字段约束
// @author: aslanwang
// @time: 2021-11-19
package main

import (
	"fmt"
	
	"github.com/go-playground/validator/v10"
)
type RegisterForm struct {
  Name      string `validate:"min=2"`
  Age       int    `validate:"min=18"`
  Password  string `validate:"min=10"`
  Password2 string `validate:"eqfield=Password"`
}

func main() {
	validate := validator.New()

  	f1 := RegisterForm{
		Name:      "dj",
		Age:       18,
		Password:  "1234567890",
		Password2: "1234567890",
	}
	err := validate.Struct(f1)
	if err != nil {
		fmt.Println(err)
	}

	f2 := RegisterForm{
		Name:      "dj",
		Age:       18,
		Password:  "1234567890",
		Password2: "123",
	}
	err = validate.Struct(f2)
	if err != nil {
		fmt.Println(err)
	}
}
```

​	结果：`f2`报错，`Key: 'RegisterForm.Password2' Error:Field validation for 'Password2' failed on the 'eqfield' tag`

## 常用约束

官方文档：[[go-playground/validator: :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving (github.com)](https://github.com/go-playground/validator)]:

知乎回答：[[学会使用validator库，看这一篇就够用了 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/194319694)]:

## gin中的参数校验

### 官方约束

​	在`validator`中，我们直接在结构体字段中将约束放到`validate tag`中，在`gin`中则只需将约束放到`binding tag`中。

### 自定义字段约束

​	示例如下：

```go
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
```

### 测试

- `2021/11/19`发送请求`http://127.0.0.1:8080/time?create_time=2021-11-20&update_time=2021-11-20`，结果如下：

![image-20211119214730625](C:\Users\wangwei\AppData\Roaming\Typora\typora-user-images\image-20211119214730625.png)

- ​	`2021/11/19`发送请求`http://127.0.0.1:8080/time?create_time=2021-11-18&update_time=2021-11-18`，结果如下：

![image-20211119214813433](C:\Users\wangwei\AppData\Roaming\Typora\typora-user-images\image-20211119214813433.png)

