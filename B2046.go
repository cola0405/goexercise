package main

import (
	"fmt"
)

func main() {
	var s float64
	fmt.Scan(&s)

	walk_time := s / 1.2
	bike_time := 27 + 23 + s/3

	if walk_time < bike_time {
		fmt.Println("Walk")
	} else if bike_time < walk_time {
		fmt.Println("Bike")
	} else {
		fmt.Println("All")
	}
}
