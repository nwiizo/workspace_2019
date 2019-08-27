package main

import (
	"fmt"
)

func main() {
	var (
		a     string
		count = 0
	)
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			count++
		}
	}
	fmt.Println(count)
}
