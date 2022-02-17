package main

import "fmt"

func main() {
	// 声明两个二维整型数组 [2]int [2]int
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2

	for index, value := range array1 {
		fmt.Printf("索引：%d,值：%d \n", index, value)
	}
}
