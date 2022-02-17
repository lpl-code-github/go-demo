package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"

	// 以下编译错误：无法比较 [2]int == [3]int
	// d := [3]int{1, 2}
	//fmt.Println(a == d)
}
