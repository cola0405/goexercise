// printf 输出百分号要%%

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := float64(b) / float64(a) * float64(100)
	fmt.Printf("%.3f%%", c)
}
