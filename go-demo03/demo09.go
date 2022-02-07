package main

import "fmt"

func main() {
	a := 5.0
	b := int(a)
	fmt.Printf("%T,%d", b, b)
}
