package main

import (
	"fmt"
)

func main() {
	var candy [5]int
	var ans int
	fmt.Scan(&candy[0], &candy[1], &candy[2], &candy[3], &candy[4])
	ans = 0

	for i := 0; i < 5; i++ {
		if candy[i]%3 != 0 {
			eaten := candy[i] % 3
			ans += eaten
		}
		amount := candy[(i+5)%5] / 3
		candy[i] = amount
		candy[(i+5-1)%5] += amount
		candy[(i+5+1)%5] += amount
	}

	fmt.Printf("%d %d %d %d %d\n", candy[0], candy[1], candy[2], candy[3], candy[4])
	fmt.Println(ans)
}
