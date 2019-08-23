package main

import "fmt" 

func Reverse(s string) string {
	runes := []rune(s)
	var i,j int 
	lens := len(runes)-1
	for i, j = 0, lens; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	s := "stressed"

	fmt.Printf("%v\n", s)
	fmt.Println(Reverse(s))

}
