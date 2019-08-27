package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkEven(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			return false
		}
	}
	return true
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	nextLine()
	inputs := strings.Split(nextLine(), " ")
	var nums []int
	for i := 0; i < len(inputs); i++ {
		num, _ := strconv.Atoi(inputs[i])
		nums = append(nums, num)
	}
	count := 0
	for checkEven(nums) {
		count++
		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] / 2
		}
	}
	fmt.Println(count)
}
