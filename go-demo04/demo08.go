package main

import "fmt"

func main() {
	var str *string = new(string)
	*str = "Go语言"
	fmt.Println(*str)
}
