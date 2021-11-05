// @program: 数组使用
// @author: aslanwang
// @create: 2021-11-5
package main

import "fmt"

// arrayUseMethod 使用数组 
// 在go中 通过名字传递数组是进行值传递 因此要想从函数中带回修改值 必须传递地址
func arrayUseMethod(){
	// 声明默认值数组
	var arr1 [5]int
	fmt.Println(arr1)
	// 声明数组并给定初始值
	var arr2 = [5]int{1, 3, 5, 7, 9}
	fmt.Println(arr2)
	// 声明数组并给定初始值 自动判断类型 这里隐含要求所有的值类型都一致 否则报错
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr3)
	// 声明数组并给定部分初始值
	arr4 := [...]int{0 : 1, 1 : 10, 4 : 68}
	fmt.Println(arr4)

	modifyArray(arr4)
	fmt.Println(arr4[4])

	modifyArrayWithAddress(&arr4)
	fmt.Println(arr4[4])
	// 直接用数组赋值给另一个数组 必须两者类型一致 大小一致
	var arr5 [5]int = arr4
	fmt.Println(arr5)
}

// modifyArray 直接传递数组是进行传值传递 同时要指明数组size 
func modifyArray(arr [5]int){
	fmt.Println(len(arr), cap(arr))
	// 在这里修改arr并不能影响原数组
	arr[4] = 498
}

// modifyArrayWithAddress 传递数组值的地址可以将结果带回原数组 
func modifyArrayWithAddress(arr *[5]int){
	fmt.Println(len(arr), cap(arr))
	// 在这里修改arr就可以影响原数组
	arr[4] = 541
}
func main(){
	arrayUseMethod()
}


 