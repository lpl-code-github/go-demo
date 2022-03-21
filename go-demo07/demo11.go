package main

import "fmt"

func main() {
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers[4:6] //{ }

	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice)) //打印出来长度为2
	fmt.Printf("myslice容量为 %d\n", cap(myslice))                  // 容量为6
	myslice = myslice[:cap(myslice)]
	// 虽然myslice 的长度为2，但容量为6，能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
