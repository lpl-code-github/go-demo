package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	slice3 := []int{7, 8}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Printf("slice2:%d\n", slice2)
	copy(slice1, slice3) // 只会复制slice3的2个元素到slice1的前2个位置
	fmt.Printf("slice1:%d", slice1)
}
