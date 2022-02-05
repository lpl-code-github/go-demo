package main

import "fmt"

func main() {
	var str = "java\ngolang"
	fmt.Println(str)

	//反引号一般用在 需要将内容进行原样输出的时候 使用
	fmt.Println(`java\ngolang`)
	fmt.Println(`\t java
golang`) //使用反引号 可以进行字符串换行

}
