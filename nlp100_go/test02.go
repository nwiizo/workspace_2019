package main

import "fmt"

func Altjoin(s1, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var s3 string
	for i := 0; i < len(runes1); i = i + 1 {
		s3 += string(runes1[i]) + string(runes2[i])
	}
	return s3
}

func main() {

	s1 := "パトカー"
	s2 := "タクシー"

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Println(Altjoin(s1, s2))
}
