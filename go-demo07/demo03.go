package main

import "fmt"

func main() {
	// 给索引为2的赋值 ，所以结果是 0,0,3
	arr := [3]int{2: 3}
	for index, value := range arr {
		fmt.Printf("索引:%d,值：%d \n", index, value)
	}
}
