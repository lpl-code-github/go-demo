package main

import "fmt"

func main() {
	// 默认数组中的值是类型的默认值
	var arr [3]int

	// 通过索引下标取值
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	println("*******")

	// for range获取
	for index, value := range arr {
		fmt.Printf("索引：%d,值：%d \n", index, value)
	}
}
