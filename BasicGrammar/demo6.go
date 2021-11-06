// @program: channel使用 
// @author: aslanwang
// @create: 2021-11-5
package main

import (
	"fmt"
	"time"
)

// channelUse 使用channel
// 1. Go 语言的协程是并发机制，交替执行。 
// 2. 声明方式
//		不带缓冲的channel 	ch1 := make(chan string)		声明带10个缓冲的通道	ch2 := make(chan string, 10)
// 		声明只读通道	ch3 := make(<-chan string)			声明只写通道	ch4 := make(chan<- string)
// 3. 声明不带缓冲的channel进行读写操作都会被阻塞	带缓冲的通道，进一次队列长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。
// 4. 关闭chan 使用close ; close 以后不能再写入，写入会出现 panic; 重复 close 会出现 panic; 只读的 chan 不能 close ;close 以后还可以读取数据
func channelUse(){
	// ch1 := make(chan string)
	// ch1 <- "Tom"
	// v := <- ch1
	ch2 := make(chan string, 3)
	// 开启一个写协程
	go func(){
		ch2 <- "hello"
	}()
	// 开启一个读协程
	go func(){
		val := <- ch2
		fmt.Println(val) 
	}()
	// 开启一个生产者
	go Producer(ch2)
	go Customer(ch2)
	time.Sleep(1 * time.Second)
}

// Producer 生产者
func Producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func Customer(ch chan string) {
	fmt.Println("customer start")
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println("customer end")
}
func main(){
	channelUse()
	fmt.Println()
}


 