package main

import (
	"fmt"
)

func sumOfDegit(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var n, a, b, sum int
	fmt.Scan(&n, &a, &b)
	for i := 1; i < n+1; i++ {
		if v := sumOfDegit(i); a <= v && v <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
