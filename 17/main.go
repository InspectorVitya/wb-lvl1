// Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

func binarySearch(a []int, search int) int {
	mid := len(a) / 2
	var index int
	switch {
	case len(a) == 0:
		// если после деления среза пополам его длина равна нулю, то искомый элемент не найден в срезе
		index = -1
	case a[mid] > search: // если искомый элемент меньше, чем средний (mid), ищем в срезе, слева от среднего (mid) элемента
		index = binarySearch(a[:mid], search)
	case a[mid] < search: // аналогично, если больше - ищем справа
		index = binarySearch(a[mid+1:], search)
		// если в результате рекурсивных вычислений вернулся положительный индекс (т.е. элемент был найден)
		// то к текущей позиции mid (средний элемент) добавляем позиции среднего элемента в срезах из рекурсивных вызовов
		if index >= 0 {
			index += mid + 1
		}
	default: // a[mid] == search
		index = mid
	}

	return index
}

func main() {
	// отсортированный срез
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	findNumber := 2
	result := binarySearch(searchField, findNumber)
	if result < 0 {
		fmt.Printf("Your number %d is not in the list\n", findNumber)
	} else {
		fmt.Printf("Your number was found in position %d\n", result)
	}
}
