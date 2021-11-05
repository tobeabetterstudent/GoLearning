// @program: 切片使用
// @author: aslanwang
// @create: 2021-11-5
package main

import "fmt"

// sliceUseMethod 使用切片 切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。 
// 在go中 通过名字传递切片是进行地址传递
func sliceUseMethod(){
	// 声明切片 nil
	var sli1 []int
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli1),cap(sli1),sli1)
	// 声明空切片
	var sli2 = []int{}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)
	// 声明切片并给定初始值 len(sli3) = cap(sli3)
	var sli3 = []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli3),cap(sli3),sli3)
	// 声明切片并给定部分初始值
	var sli4 []int = make([]int,5,8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli4),cap(sli4),sli4)

	modifyslice(sli4)
	fmt.Println("After updated sli[4]:", sli4[4])
}

// interceptSlice 截取切片
func interceptSlice(){
	sli := [] int {1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	fmt.Println("sli[1] ==", sli[1])
	// 截取切片的所有元素
	fmt.Println("sli[:] ==", sli[:])
	// 从切片的index = 1处截取切片 包括index = 1
	fmt.Println("sli[1:] ==", sli[1:])
	// 从切片的开始截取到index = 4 不包括index = 4 len = 4
	fmt.Println("sli[:4] ==", sli[:4])
	// 从切片的index = 0 截取到index = 3 不包括index = 3 len = 3 - 0
	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3]),cap(sli[0:3]),sli[0:3])
	// 从切片的index = 0 截取到index = 3 不包括index = 3 len = 3 - 0 并且指定 cap = 4
	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3:4]),cap(sli[0:3:4]),sli[0:3:4])
}

// appendSlice 追加切片 在调用append 时，容量不够需要扩容时，cap 会翻倍。
func appendSlice(){
	sli := [] int {4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}

// deleteSlice 删除切片的某些元素 巧妙使用append
func deleteSlice(){
	sli := [] int {1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	//删除尾部 2 个元素 sli[:len(sli) - 2]截取的元素长度为len(sli) - 2
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[:len(sli)-2]),cap(sli[:len(sli)-2]),sli[:len(sli)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[2:]),cap(sli[2:]),sli[2:])

	//删除中间 2 个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}
// modifyslice 直接传递切片是进行地址传递
func modifyslice(sli []int){
	fmt.Println(len(sli), cap(sli))
	// 在这里修改sli并不能影响原切片
	sli[4] = 498
}

func main(){
	sliceUseMethod()
	fmt.Println()
	interceptSlice()
	fmt.Println()
	appendSlice()
	fmt.Println()
	deleteSlice()
}


 