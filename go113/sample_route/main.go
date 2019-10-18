package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	for i := 0; i < 5; i++ {
		go func(i int) {
			// iを参照し続ける
			Greet(i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("main end")
}
func Greet(i int) {
	fmt.Println("hello", i)
}
