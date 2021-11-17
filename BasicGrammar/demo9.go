// @program: 错误处理
// @author: aslanwang
// @create: 2021-11-17
package main

import (
	"errors"
	"fmt"
)

// Hello 
// 在函数声明时已指定 str 和 err 是返回值，相当于在函数第一句定义该变量，在遇到 return 语句时它们会自动从函数中返回。
// 在 Golang 中，有返回值的函数，无论是命名返回值还是普通形式的返回值，函数中必须包含 return 语句。
func Hello(name string) (str string,err error){
	if name == "" {
		err = errors.New("Name Is Empty!")
		return
	}
	str = fmt.Sprintf("Hello %s", name)
	return
}

func main(){
	var name string
	fmt.Scanln(&name)
	str, err := Hello(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)
}