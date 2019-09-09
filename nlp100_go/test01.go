package main

import "fmt"

func oddjoin(s1 string) string {
	runes1 := []rune(s1) //rune配列に変換
	var s2 string
	for i := 0; i < len(runes1); i = i + 1 {
		if i % 2 == 1{
			s2 += string(runes1[i])
		}
	}
	return s2
}

func main() {

	s := "パタトクカシーー"

	fmt.Printf("%v\n", s)
	fmt.Println(oddjoin(s))

}
