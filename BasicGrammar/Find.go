// @program: 不定参数传递
// @author: aslanwang
// @create: 2021-11-6
package main

import (
	"BasicGrammar/friend"
	"fmt"
)

func main() {
	friends, err := friend.Find("附近的人",
		friend.WithSex(1),
		friend.WithAge(30),
		friend.WithHeight(160),
		friend.WithWeight(55),
		friend.WithHobby("爬山"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(friends)
}