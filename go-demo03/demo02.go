package main

import "fmt"

func main() {
	var aVar = 10
	println(aVar == 5)  // false
	println(aVar == 10) // true
	println(aVar != 5)  // true
	println(aVar != 10) // false

	var a = 10
	//因为a>11已经不满足了，所以a < 30不会执行，整个表达式为false
	if a > 11 && a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}

	//因为a > 5已经满足了，所以a < 30不会执行，整个表达式为true
	if a > 5 || a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}
}
