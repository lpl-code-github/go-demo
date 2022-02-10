package main

import (
	"fmt"
)

func main() {
	var cat int = 1
	var str string = "go语言"
	fmt.Printf("%p %p", &cat, &str)
}
