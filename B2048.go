package main

import (
	"fmt"
)

func main() {
	var w int
	var s string
	fmt.Scan(&w, &s)

	price := 8
	if w > 1000 {
		price += (w - 1000) / 500
	}
}
