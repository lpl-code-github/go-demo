package main

import (
	"fmt"
	"strconv"
)

func main() {
	newStr := "3.1415926"
	floatValue, _ := strconv.ParseFloat(newStr, 32)
	fmt.Printf("floatValue的类型：%T, floatValue的值：%f", floatValue, floatValue)
}
