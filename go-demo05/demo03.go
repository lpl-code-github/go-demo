package main

import "fmt"

/*
	算数运算符
*/
func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("c = a + b， c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("c = a - b， c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("c = a * b， c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("c = a / b， c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("c = a ÷ b， c 的值为 %d\n", c)
	a++
	fmt.Printf("a++， a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("a--， a 的值为 %d\n", a)
}
