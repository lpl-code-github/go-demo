package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//如何计算字符串的长度
	str3 := "hello"
	str4 := "你好"
	fmt.Println(len(str3))                    // 1个字母占1个字节
	fmt.Println(len(str4))                    // 1个中文占3个字节，go从底层支持utf8
	fmt.Println(utf8.RuneCountInString(str4)) // 2

}
