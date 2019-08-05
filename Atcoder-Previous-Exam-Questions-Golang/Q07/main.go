package main

import (
	"fmt"
)

func main() {
	var n, mochi int
	fmt.Scan(&n)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&mochi)
		m[mochi]++
	}

	fmt.Println(len(m))
}
