package main

import "fmt"

const (
	Sunday    = iota //0
	Monday           // +1 = 1
	Tuesday          // +1 = 3
	Wednesday        // +1 = 4
	Thursday         // +1 = 5
	Friday           // +1 = 6
	Saturday         // +1 = 7
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
