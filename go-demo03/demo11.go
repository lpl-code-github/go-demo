package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello,Java语言"
	source := "Java"
	target := "Go"
	index := strings.Index(str1, source) // "Java"的位置 6
	sourceLength := len(source)          // "Java"的长度 4
	start := str1[:index]                // 第一部分切片"hello,"
	end := str1[index+sourceLength:]     // 第二部分切片"语言"

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(start)
	stringBuilder.WriteString(target)
	stringBuilder.WriteString(end)
	fmt.Println(stringBuilder.String())

}
