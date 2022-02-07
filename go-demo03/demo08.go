package main

import (
	"fmt"
	"strings"
)

func main() {
	tracer := "你好世界,再见世界"

	// 正向搜索字符串
	str1 := strings.Index(tracer, "世")
	fmt.Println("世第一次出现的位置:", str1)

	// 反向搜索字符串
	str2 := strings.LastIndex(tracer, "世")
	fmt.Println("世最后一次出现的位置:", str2)
}
