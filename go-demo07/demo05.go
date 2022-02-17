package main

import "fmt"

func main() {
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	fmt.Println(array[1][0])

	println("***********")

	for index, value := range array {
		fmt.Printf("索引:%d,值：%d \n", index, value)
	}
}
