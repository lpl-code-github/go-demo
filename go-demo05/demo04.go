package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("a 等于 b\n")
	} else {
		fmt.Printf("a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("a 小于 b\n")
	} else {
		fmt.Printf("a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("a 大于 b\n")
	} else {
		fmt.Printf("a 不大于 b\n")
	}

	a = 5
	b = 20
	if a <= b {
		fmt.Printf("a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("b 大于等于 a\n")
	}
}
