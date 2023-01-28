// 参与运算的是什么类型，得到的就是什么类型
package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	if num1/10 >= 1 || num2/20 >= 1 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
