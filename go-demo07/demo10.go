package main

import "fmt"

func main() {
	// make( []Type, size, cap )
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	//容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2
	//但如果我们给a 追加一个 a的长度就会变为3
	fmt.Println(len(a), len(b))

	b = append(b, 1)
	fmt.Println(len(a), len(b))
}
