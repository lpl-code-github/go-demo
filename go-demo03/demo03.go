package main

import "fmt"

func main() {
	//使用单引号 表示一个字符
	var ch1 byte = 'A'
	//在 ASCII 码表中，A 的值是 65,也可以这么定义
	var ch2 byte = 65
	//65使用十六进制表示是41，所以也可以这么定义 \x 总是紧跟着长度为 2 的 16 进制数
	var ch3 byte = '\x41'
	//65的八进制表示是101，所以使用八进制定义 \后面紧跟着长度为 3 的八进制数
	var ch4 byte = '\101'

	fmt.Printf("%c,", ch1)
	fmt.Printf("%c,", ch2)
	fmt.Printf("%c,", ch3)
	fmt.Printf("%c", ch4)
}
