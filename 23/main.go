// Удалить i-ый элемент из слайса.
package main

import "fmt"

func remove(arr []int, i int) ([]int, error) {
	if i >= 0 && i < len(arr) {
		res := make([]int, 0, len(arr)-2)
		res = append(res, arr[:i]...)
		return append(res, arr[i+1:]...), nil // arr[:i+copy(arr[i:], arr[i+1:])]
	}
	return []int{}, fmt.Errorf("invalid index")
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr, _ = remove(arr, 4)
	fmt.Println(arr)
}
