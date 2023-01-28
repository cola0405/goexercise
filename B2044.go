package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var count int
	count = 0
	if a < 60 {
		count++
	}
	if b < 60 {
		count++
	}
	if c < 60 {
		count++
	}
	if count == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
