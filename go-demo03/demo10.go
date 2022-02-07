package main

import "fmt"

func main() {

	s1 := "localhost:8080"
	fmt.Println(s1)
	// 强制类型转换 string to byte
	strByte := []byte(s1)

	// 遍历
	for _, s := range strByte {
		fmt.Printf("strByte: %c\n", s)
	}

	// 下标修改
	strByte[len(s1)-1] = '1'

	// 强制类型转换 []byte to string
	s2 := string(strByte)
	fmt.Println(s2)
}
