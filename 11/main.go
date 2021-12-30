/*
	Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
	out := make([]int, 0)
	m := make(map[int]bool)

	for _, v := range arr1 {
		m[v] = true
	}
	for _, val := range arr2 {
		if m[val] {
			out = append(out, val)
		}
	}
	return out
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 9, 10}
	set2 := []int{2, 4, 8, 10, 50, 5, 8, 9, 10}
	fmt.Println(intersection(set1, set2))
}
