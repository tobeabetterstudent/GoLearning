// @program: 结构体使用 与 json转化 
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

// structUse 声明struct并进行json化
func structUse(){
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

// strongTypeJson 强类型json转struct
// 在线工具: https://mholt.github.io/json-to-go/
type MobileInfo struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Areacode string `json:"areacode"`
		Zip      string `json:"zip"`
		Company  string `json:"company"`
		Card     string `json:"card"`
	} `json:"result"`
}

func strongTypeJson(){
	jsonStr := `{
		"resultcode": "200", "reason": "Return Successd!",
		"result": {
			"province": "浙江", "city": "杭州",
			"areacode": "0571", "zip": "310000",
			"company": "中国移动", "card": ""
			}
		}`
	
	var mobile MobileInfo
	err := json.Unmarshal([]byte(jsonStr), &mobile)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(mobile)
}

func main(){
	structUse()
	fmt.Println()
	strongTypeJson()
}


 