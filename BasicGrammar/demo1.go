// @program: 变量声明
// @author: aslanwang
// @create: 2021-11-5
package main

import "fmt"

// constantDeclaration 常量声明 2种方式
// 1 const arg1, arg2... type = value1, value2, ...
// 2.const arg1, arg2... = value1, value2, ... 自行判断数据类型

func constantDeclaration(){
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name1, name2 string = "Tom", "Alice"
	fmt.Println(name1, name2)

	const age1, age2 = 40, 50
	fmt.Println(age1, age2)
}

// varabileDeclaration 变量声明 3种方式
// 1 var arg1, arg2... type = value1, value2, ...
// 2.var arg1, arg2... = value1, value2, ...自行判断数据类型
// 3. arg1, arg2 ... = value1, value2, ... 	自行判断数据类型

func varabileDeclaration(){
	var age1 uint8 = 6
	fmt.Println(age1)
	var age2 = 8
	fmt.Println(age2)
	age3 := 10
	fmt.Println(age3)

	var t1, t2, t3 uint8 = 1, 3, 5
	fmt.Println(t1, t2, t3)
	var a1, a2, a3 = 7, "world", 8.4
	fmt.Println(a1, a2, a3)
	v1, v2, v3 := 7 , "hello", 5.6
	fmt.Println(v1, v2 ,v3)
}

func main(){
	constantDeclaration()
	varabileDeclaration()
	fmt.Printf("name = %s, age = %d, height = %v", "Tom", 18, 165.67)
}


 