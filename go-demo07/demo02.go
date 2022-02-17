package main

import "fmt"

func main() {
	// 1. 初始化赋值
	var arr01 [3]int = [3]int{1, 2, 3}
	// 1.1 如果第三个不赋值，就是默认值0
	//var arr01 [3]int = [3]int{1, 2}
	// 1.2 可以使用简短声明
	//arr01 := [3]int{1, 2, 3}
	// 1.3 如果不写数据数量，而使用...，表示数组的长度是根据初始化值的个数来计算
	//arr01 := [...]int{1, 2, 3}

	// 2. 通过索引下标赋值
	var arr02 [3]int
	arr02[0] = 5
	arr02[1] = 6
	arr02[2] = 7

	for index, value := range arr01 {
		fmt.Printf("arr01-索引：%d,值：%d \n", index, value)
	}

	println("****************")

	for index, value := range arr02 {
		fmt.Printf("arr02-索引：%d,值：%d \n", index, value)
	}

}
