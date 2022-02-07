package main

import (
	"bytes"
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

	// 字符串拼接
	s := "第一部分" + "第二部分"
	fmt.Println(s)

	str := "hel" + "lo,"
	str += "world!"
	fmt.Println(str) //输出 “hello, world!”

	str1 := "你好,"
	str2 := "世界"
	var stringBuilder bytes.Buffer
	//节省内存分配，提高处理效率
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	fmt.Println(stringBuilder.String())

	// 从字符串 hello 世界 中获取 世
	var myStr01 string = "hello,世界"
	fmt.Println(string([]rune(myStr01)[6]))
}
