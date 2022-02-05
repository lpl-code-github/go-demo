package main

import "fmt"

func main() {
	var str = "java\ngolang"
	fmt.Println(str)

	//反引号一般用在 需要将内容进行原样输出的时候 使用
	fmt.Println(`java\ngolang`)
	fmt.Println(`\t java
golang`) //使用反引号 可以进行字符串换行

	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("myStr01: %s\n", mystr01)
	fmt.Printf("myStr02: %s\n", mystr02)

	var myStr01 string = "hello,世界"
	fmt.Printf("myStr01: %d\n", len(myStr01)) //输出结果 myStr01: 12
}
