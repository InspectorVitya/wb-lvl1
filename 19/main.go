/*
	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/
package main

import "fmt"

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	question := "nwòrb"
	fmt.Printf("%s\n%s", question, ReverseString(question))
}
