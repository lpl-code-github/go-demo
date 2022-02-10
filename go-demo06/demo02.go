package main

import (
	"fmt"
	"strconv"
)

func main() {
	intValue2 := 1
	strValue := strconv.Itoa(intValue2)
	fmt.Printf("strValue类型:%T, strValue值：%s\n", strValue, strValue)
}
