package main

import "fmt"

func main() {
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据 切片不会因为等号操作进行元素的复制
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素和最后一个元素
	srcData[0], srcData[elementCount-1] = 777, 500
	// 打印引用切片的第一个元素 引用数据的第一个元素将会发生变化
	fmt.Printf("引用切片的第一个元素:%d,最后一个元素:%d\n", refData[0], refData[elementCount-1])
	// 打印复制切片的第一个和最后一个元素 由于数据是复制的，因此不会发生变化。
	fmt.Printf("复制切片的第一个元素:%d,最后一个元素:%d\n", copyData[0], copyData[elementCount-1])
}
