/* Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (02,04,06,08,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// Версия с запуском возведение в квадрат в отдельной горутине.
// Ожидаем выполнение всех горутин при помощи WaitGroup и возвращаем заполненный слайс.
func powWg(in []int) []int {
	powArr := make([]int, len(in))
	var wg sync.WaitGroup
	for index, value := range in {
		wg.Add(1)
		go func(value int, res []int, index int) {
			defer wg.Done()
			powArr[index] = value * value
		}(value, powArr, index)
	}
	wg.Wait()
	return powArr
}

// Вариант с функцией (будем запускать её в горутине), пишущей квадраты в канал
// и закрывающей его по окончанию записи (элементов массива)
// в качестве параметра принимает chan<- - указание, что канал только для записи
func powChan(c chan<- int, in []int) {
	for _, v := range in {
		c <- v * v
	}
	close(c)
}

func main() {
	inputArr := []int{2, 4, 6, 8, 10}

	res := powWg(inputArr)
	fmt.Println(res)

	channel := make(chan int)
	go powChan(channel, inputArr)

	for pow := range channel {
		fmt.Println(pow)
	}
}
