package main

import "fmt"

func main() {
	const (
		i = 1 << iota //1<<0  二进制 0000 0001
		j = 3 << iota //3<<1  二进制 0000 0110
		k             //3<<2  二进制 0000 1100
		l             //3<<3  二进制 0001 1000
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
