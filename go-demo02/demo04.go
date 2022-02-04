package main

import (
	"fmt"
)

//全局变量 x
var x int = 13

func main() {
	//局部变量 x 和 y
	var x int = 3
	var y int = 4
	fmt.Printf("main() 函数中 x = %d\n", x)
	fmt.Printf("main() 函数中 y = %d\n", y)
	c := sum(x, y)
	fmt.Printf("main() 函数中 c = %d\n", c)
}
func sum(x, y int) int {
	fmt.Printf("sum() 函数中 x = %d\n", x)
	fmt.Printf("sum() 函数中 y = %d\n", y)
	num := x + y
	return num
}
