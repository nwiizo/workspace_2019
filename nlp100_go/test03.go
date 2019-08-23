package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	s = strings.Trim(s, ".")
	s = strings.Replace(s, ",", "", -1)

	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		fmt.Printf("%d:", len(words[i]))
		fmt.Printf("%v\n", words[i])
	}

}
