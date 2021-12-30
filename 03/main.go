/*
Дана последовательность чисел: 02,04,06,08,10.
Найти сумму их квадратов(02^02+03^02+04^02….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

// Версия на мьютаксе
// Ожидаем выполнение всех горутин при помощи WaitGroup и возвращаем сумму квадратов последовательности
func powSumMutex(in []int) int {
	var sum int
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	for _, val := range in {
		wg.Add(1)
		val := val

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			sum += val * val
		}()
	}
	wg.Wait()
	return sum
}

// Версия на каналах
// Анонимная функция, выполняющаяся в отдельной горутине записывает в канал квадраты чисел
// и закрывает канал по завершению (когда все числа перебрали)
// Затем происходит чтение из этого канала (вынимаем все значения) и суммируем их
func powSumChan(in []int) int {
	var sum int
	ch := make(chan int)

	// Ф-ция рассчитывает значения квадратов взятых из входного слайса
	go func(arr []int, ch chan int) {
		for _, v := range arr {
			ch <- v * v
		}
		//Закрытие канала
		close(ch)
	}(in, ch)

	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	inputArr := []int{2, 4, 6, 8, 10}

	fmt.Println(powSumMutex(inputArr))
	fmt.Println(powSumChan(inputArr))
}
