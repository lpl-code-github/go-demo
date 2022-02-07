package main

import "fmt"

func main() {
	//unicode字符集使用`for range`进行遍历
	//ascii字符集可以使用`for range`或者`for`循环遍历
	var str1 string = "hello"
	var str2 string = "hello,世界"
	// 遍历
	for i := 0; i < len(str1); i++ {
		fmt.Printf("for循环遍历str1: %c %d\n", str1[i], str1[i])
	}
	for _, s := range str1 {
		fmt.Printf("for range遍历str1: %c %d\n", s, s)
	}
	// 中文只能用 for range
	for _, s := range str2 {
		fmt.Printf("for range遍历str2: %c %d\n", s, s)
	}
}
