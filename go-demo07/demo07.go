package main

import "fmt"

//func main() {
//	var a = [3]int{1, 2, 3}
//	//a[1:len(a)] 生成了一个新的切片
//	fmt.Printf("切片前:%d \n切片后:%d", a, a[1:len(a)])
//}

func main() {
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
}
