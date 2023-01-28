// 保留小数是四舍五入

package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	d := 2 * r
	s := math.Pi * math.Pow(r, float64(2))
	p := 2 * math.Pi * r
	fmt.Printf("%.4f %.4f %.4f", d, p, s)

}
