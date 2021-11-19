package entity

// 定义 Member 结构体
// 其中tag:binding的required gt lt 是已有字段 而validate标签是自定义的 包括其后面的参数校验方法NameVaild
type Member struct {
	Name string `form:"name" json:"name" binding:"required,NameVaild"`
	Age  int    `form:"age"  json:"age"  binding:"required,gt=10,lt=120"`
}