package main

import "fmt"

func main() {
	// 批量声明多个常量
	const (
		//如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略
		a = 1
		//如果省略初始化表达式则表示使用前面常量的初始化表达式，对应的常量类型也是一样的
		b // 此时b 也会初始化为1
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
}
