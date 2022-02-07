package main

import (
	"bytes"
	"fmt"
)

func main() {
	str1 := "你好,"
	str2 := "世界"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	// Sprint 以字符串形式返回
	result := fmt.Sprintf(stringBuilder.String())
	fmt.Println(result)
}
