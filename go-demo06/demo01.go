package main

import (
	"fmt"
	"strconv"
)

func main() {
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("%T,%d\n", intValue, intValue)
}
