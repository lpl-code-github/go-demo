package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "运行模式，可以设置为fast")

func main() {
	// 解析命令行参数
	flag.Parse()
	fmt.Printf("运行模式为:%s", *mode)
}
