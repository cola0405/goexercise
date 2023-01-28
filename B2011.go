package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var c float64
	c = float64(a) / float64(b)
	fmt.Printf("%.9f", c)
}
