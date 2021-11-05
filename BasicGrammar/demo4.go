// @program: 结构体使用
// @author: aslanwang
// @create: 2021-11-5
package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	Name string `json:"name"`
	Age int		`json:"age"`		// 如果将之改为 age 则该字段在json时不可见
}

func main(){
	p1 := new(person)
	p1.Name = "Tom"
	p1.Age = 18
	fmt.Println(p1)
	data,err := json.Marshal(p1)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Println(string(data))
	var p2 person
	err = json.Unmarshal(data, &p2)
	fmt.Println(p2)

}


 