package main

import "fmt"

//声明全局变量
var c1 int

//声明全局变量
var bb float32 = 3.14

func main() {
	//声明局部变量
	var a, b int
	//初始化参数
	a = 3
	b = 4
	c1 = a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c1)

	bb := 3
	fmt.Println(bb)
}
