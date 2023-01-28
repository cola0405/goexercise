package main

import (
	"fmt"
	"math"
)

func main() {
	var h, r float64
	fmt.Scan(&h, &r)
	v := float64(h) * math.Pi * math.Pow(r, 2)
	n := float64(20*1000) / v
	fmt.Println(int(math.Ceil(n)))
}
