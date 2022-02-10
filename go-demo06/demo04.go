package main

import (
	"fmt"
	"strconv"
)

func main() {
	floatValue := 3.1415926
	/*
		4个参数:
			1、要转换的浮点数
			2、格式标记（b、e、E、f、g、G）
				‘b’ (-ddddp±ddd，二进制指数)
				‘e’ (-d.dddde±dd，十进制指数)
				‘E’ (-d.ddddE±dd，十进制指数)
				‘f’ (-ddd.dddd，没有指数)
				‘g’ (‘e’:大指数，‘f’:其它情况)
				‘G’ (‘E’:大指数，‘f’:其它情况)
				如果格式标记为 ‘e’，‘E’和’f’，则 prec 表示小数点后的数字位数
				如果格式标记为 ‘g’，‘G’，则 prec 表示总的数字位数（整数部分+小数部分）
			3、精度
			4、指定浮点类型（32:float32、64:float64）
	*/
	formatFloat := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Printf("formatFloat的类型%T，formatFloat的值%s", formatFloat, formatFloat)
}
