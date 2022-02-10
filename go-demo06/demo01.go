package main

import (
	"fmt"
	"strconv"
)

func main() {
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("ingValue类型:%T，intValue值:%d", intValue, intValue)
}
