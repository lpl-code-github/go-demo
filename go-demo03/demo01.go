package main

import "fmt"

// 浮点型
func main() {
	//floatNum := 3.2
	//// 保留小数点后2位数
	//fmt.Printf("%.2f\n", floatNum)

	// float32 类型的累计计算误差很容易扩散
	var f float32 = 1 << 24
	fmt.Println(f)
	fmt.Printf("%.0f,%.0f", f, f+1)
	fmt.Println("")
	fmt.Println(f == f+1)

	// 浮点数在声明的时候可以只写整数部分或者小数部分
	var e = .71828
	var g = 1.
	fmt.Printf("%.5f,%.1f", e, g) // 执行结果 0.71828,1.0

	// 很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分
	var avogadro = 6.02214129e23 // 阿伏伽德罗常数
	var planck = 6.62606957e-34  // 普朗克常数
	fmt.Printf("%f,%.35f", avogadro, planck)
}
