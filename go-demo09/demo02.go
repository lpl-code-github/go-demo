package main

import "fmt"

func main() {
	scene := make(map[string]int, 3)
	scene["cat"] = 66
	scene["dog"] = 4
	scene["pig"] = 960
	// 删除键值对
	delete(scene, "dog")
	// 遍历
	for k, v := range scene {
		fmt.Printf("key:%s,value:%d\n", k, v)
	}
}
