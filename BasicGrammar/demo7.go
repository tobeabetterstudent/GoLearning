// @program: defer函数的使用 
// @author: aslanwang
// @create: 2021-11-6
package main

import (
	"fmt"
)

// deferExcuteOrder 首先探究defer函数执行的时机和顺序
// 1.当使用os.Exit()方法退出程序时，defer不会被执行
// 2.defer 只对当前协程有效
// 3.defer函数在外部函数return后执行 其执行的顺序与声明的顺序相反
// 4.对于defer调用的函数传参时参数值的说明 
// 		由defer声明的函数，参数的值在 defer 定义时就会去计算这个值了/ 相当于传值 仅仅推迟执行了函数体
//			如defer fmt.Println(a + b) 或defer func(a,b int){fmt.Println("defer func_2: a + b = ",a+b)}(a,b)
//		而defer函数内部所使用的变量的值需要在这个函数运行时才确定，如defer func() { fmt.Println(a + b) }()
func deferExcuteOrder(){
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()

	a, b := 1, 2
	defer fmt.Println("defer func_1: a + b = ",a+b)
	defer func(a,b int){
		fmt.Println("defer func_2: a + b = ",a+b)
	}(a,b)
	defer func(){
		fmt.Println("in defer func: a + b = ",a+b)
	}()
	a = 2
	fmt.Println("End------------------End")
}

// t1 输出1 
// 再返回1之后 defer才执行 但这时跟return的值 那个值与a不是一个地址
func t1() int {
	a := 1
	defer func(){
		a++;
	}()
	return a
}

// t2 输出2
// 再返回1之后 defer执行 这时函数t2还在栈中 defer函数直接操作了返回值+1 因此返回2
func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

// ----------------------------------------------------------------------
// calc 在deferUse中 使用 以说明defer 函数 时遇到的传参问题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func deferUse() {
	x := 1
	y := 2
	// 1. defer声明的 函数 在定义时就去计算值 因此x = 1  y = 2并且执行calc("B", x, y) 输出B 1 2 3
	defer calc("A", x, calc("B", x, y))
	// 4. 在函数退出后 defer函数推迟的函数体按从下向上执行 执行calc("A", x, 3) 输出 A 1 3 4
	x = 3
	// 2. defer声明的 函数 在定义时就去计算值 因此x = 3  y = 2并且执行calc("D", x, y) 输出D 3 2 5
	defer calc("C", x, calc("D", x, y))
	// 3. 在函数退出后 defer函数推迟的函数体按从下向上执行 执行calc("C", x, 5) 输出 C 3 5 8
	y = 4
}
// ----------------------------------------------------------------------
func main(){
	deferExcuteOrder()
	fmt.Println("t1 : ", t1())
	fmt.Println("t2 : ", t2())
	deferUse()
}


 