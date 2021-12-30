/*
	Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// Завершение с помощью сигнализирующего канала
func first(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("stop first goroutine")
			return
		default:
			fmt.Println("work")
		}
	}
}

// Горутина работает пока не закрыт канал
func second(ch chan int) {
	for {
		_, ok := <-ch // проверяем закрыт ли канал
		if !ok {
			fmt.Println("stop second goroutine")
			return
		}
		fmt.Println("work")
	}
}

// Горутина работает пока не закрыт канал
func third(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("stop third goroutine")
}

//	Завершение с помощью контекста
func forth(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если контекст был отменен
			fmt.Println("stop forth goroutine")
			return
		default:
			fmt.Println("work")
		}
	}
}

// Завершение горутины по истечению времени
func fifth() {
	timeOut := time.After(1 * time.Microsecond)
	for {
		select {
		case <-timeOut:
			fmt.Println("stop fifth goroutine")
			return
		default:
			fmt.Println("work")
		}
	}
}

func main() {
	done := make(chan struct{})
	go first(done)
	done <- struct{}{}

	intChan := make(chan int)
	go second(intChan)
	go third(intChan)
	close(intChan)

	ctx, cancel := context.WithCancel(context.Background())
	go forth(ctx)
	cancel()

	go fifth()

	time.Sleep(time.Second)
}
