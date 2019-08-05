package main

import (
	"fmt"
)

func checkSubstr(s string, d []string) bool {
	for i := 0; i < len(d); i++ {
		//s[0:len(d[i])]のエラーチェック
		if len(s) < len(d[i]) {
			continue
		}
		if s[0:len(d[i])] == d[i] {
			if len(s) == len(d[i]) {
				return true
			}
			return checkSubstr(s[len(d[i]):], d)
		}
	}
	return false
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	var s string
	fmt.Scan(&s)
	d := []string{"maerd", "remaerd", "esare", "resare"}
	s = reverse(s)

	if checkSubstr(s, d) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
