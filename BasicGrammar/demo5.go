// @program: map使用 与 json转化 
// @author: aslanwang
// @create: 2021-11-5
package main

import (
	"fmt"
	"encoding/json"
)

// mapUse 声明map并进行json化 map使用时必须进行初始化 否则是nil 可以通过make 或者 直接一个map赋值
// go语言的map是无序的
// interface{}可以接收任意类型的变量
// 只有map[string]Type 可以进行json序列化
func mapUse(){
	var mp1 map[string]int = map[string]int{"Tom":1, "Jenny":2}
	fmt.Println(mp1) 

	mp2 := make(map[string]int)
	mp2["Black"] = 3
	fmt.Println(mp2)

	fmt.Println()
	// key为string value可以接受任何变量
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"]  = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age"      : "30",
		"hobby"    : []string{"读书","爬山"},
	}
	fmt.Println("map data :", res)
	// 序列化
	jsons, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Println("--- res to json ---")
	fmt.Println("json data :", string(jsons))
	// 反序列化
	res2 := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsons), &res2)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("--- json to map ---")
	fmt.Println("map data :",res2)
}

func main(){
	mapUse()
	fmt.Println()
}


 