package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num >= 10 && num <= 99 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
