package main

import "fmt"

func main() {

	var a = 100
	var b = 200
	// 变量交换 第一种
	var c int
	c = b
	b = a
	a = c
	fmt.Printf("a=%d,b=%d", a, b)

}
