package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, s float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	xGap := math.Abs(x1 - x2)
	yGap := math.Abs(y1 - y2)
	s = math.Sqrt(math.Pow(xGap, 2) + math.Pow(yGap, 2))
	fmt.Printf("%.3f", s)
}
