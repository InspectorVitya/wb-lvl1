/*
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func IsUniqSymbols(str string) bool {
	m := make(map[rune]bool)
	str = strings.ToLower(str)
	for _, s := range str {
		if m[s] {
			return false
		}
		m[s] = true
	}
	return true
}

func main() {
	test := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, val := range test {
		fmt.Println(IsUniqSymbols(val))
	}
}
