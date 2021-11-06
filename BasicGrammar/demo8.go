// @program: 结构体实现接口
// @author: aslanwang
// @create: 2021-11-6
package main

import (
	"fmt"
	"errors"
)

// 强制*student类型的变量去实现Study接口 否则在编译时报错
var _ Study = (*student)(nil)

type Study interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

type student struct {
	Name string
}

func (s *student) Listen(msg string) string {
	return s.Name + " 听 " + msg
}

func (s *student) Speak(msg string) string {
	return s.Name + " 说 " + msg
}

func (s *student) Read(msg string) string {
	return s.Name + " 读 " + msg
}

func (s *student) Write(msg string) string {
	return s.Name + " 写 " + msg
}

func NewStudent(name string) (Study, error) {
	if name == "" {
		return nil, errors.New("name required")
	}

	return &student{
		Name: name,
	}, nil
}

func implInterface(){
	name := "Tom"
	s, err := NewStudent(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}

func main(){
	implInterface()
}


 