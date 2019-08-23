package main

import (
	"fmt"
	"strings" 
)

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	s = strings.Trim(s, ".")

	s = strings.Replace(s, ",", "", -1)
	words := strings.Split(s, " ") 

	words_index := make(map[string]int)

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		switch i {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19: 
			words_index[string(runes[0])] = i
		default:
			words_index[string(runes[0])+string(runes[1])] = i
		}
	}
	fmt.Println(words_index)
}
