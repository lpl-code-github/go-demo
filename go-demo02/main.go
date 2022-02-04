package main

import "fmt"

// 标准格式声明变量
var age int

// 游戏角色初始等级 不指明变量类型
var level = 1

// 批量声明变量
var (
	a int
	b string
	c float32
)

func main() {
	// %d 代表要打印一个整数类型
	//fmt.Printf("%d", age)

	// %T 打印类型
	//fmt.Printf("%T", level)

	// %d 整数占位符，%s 字符串占位符， %f 浮点数占位符(默认精度为6)
	fmt.Printf("%d,%s,%f", a, b, c)
}
