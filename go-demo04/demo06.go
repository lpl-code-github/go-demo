package main

import (
	"fmt"
)

func main() {
	var cat int = 1
	var str string = "go语言"
	ptr := &cat

	fmt.Printf("%p %p\n", &cat, &str) //变量地址
	fmt.Println(ptr)                  // 变量地址
	fmt.Println(*ptr)                 //使用指针访问值
}
