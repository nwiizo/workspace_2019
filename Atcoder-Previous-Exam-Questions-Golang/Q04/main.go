package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c, x, count int
	)
	fmt.Scan(&a, &b, &c, &x)
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if i*500+j*100+k*50 == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

