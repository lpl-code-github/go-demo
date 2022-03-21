package main

import "fmt"

func main() {
	var strList []string
	fmt.Printf("追加之前:%s\n", strList)
	// 追加一个元素
	strList = append(strList, "go语言")
	fmt.Printf("追加之后:%s", strList)
}
