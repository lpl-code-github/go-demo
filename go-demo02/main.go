package main

import "fmt"

// 标准格式声明变量
var age int

// 游戏角色初始等级 不指明变量类型
var level = 1

func main() {
	// %d 代表要打印一个整数类型
	//fmt.Printf("%d", age)

	// %T 打印类型
	fmt.Printf("%T", level)
}
