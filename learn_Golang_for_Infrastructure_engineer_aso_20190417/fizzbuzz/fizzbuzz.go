package main

import (
	"fmt"
)

func fizzbuzz(x int) (result string) {

	if x%3 == 0 && x%5 == 0 {
		result = "FizzBuzz"
	} else if x%3 == 0 {
		result = "Fizz"
	} else if x%5 == 0 {
		result = "Buzz"
	} else {
		result = ""
	}
	return

}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(i, fizzbuzz(i))
	}
}
