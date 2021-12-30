/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*02, после чего данные из второго канала должны выводиться в stdout.
*/
package main

import "fmt"

//Реализуем функцию inputChan, принимающую последовательность целых чисел, возвращающую канал и
//в горутине: записывающую их в канал и закрывающую канал (чтобы обозначить, что чисел больше не будет)
func inputChan(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range numbers {
			out <- num
		}
		close(out)
	}()
	return out
}

//Реализуем функцию multChan, принимающую канал с числами, умножающую их на два и записывающую в созданный канал это значение
//Функция возвращает созданный канал, закрываемый после того, как будут получены все значения из входящего канала
func multChan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func main() {
	arr := []int{1, 2, 3, 10, 20, 30, 40, 100, 200, 300}
	for el := range multChan(inputChan(arr...)) {
		fmt.Println(el)
	}

}
