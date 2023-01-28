package main

import (
	"fmt"
)

func main() {
	var w int
	fmt.Scan(&w)

	busy := [3]int{1, 3, 5}

	for i := range busy {
		if busy[i] == w {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
