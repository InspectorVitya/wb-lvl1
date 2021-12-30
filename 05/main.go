/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Print("Input working time(sec): ")
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalln("Bad number of time value, program is shutting down")
	}

	timeOut := time.After(time.Duration(n) * time.Second)
	ch := make(chan int)
	done := make(chan struct{})
	// Читаем из канала пока не вышло время
	// Когда выходит время, отправляем сигнал в done канал.
	go func() {
		for {
			select {
			case <-timeOut:
				done <- struct{}{}
				break
			default:
				fmt.Println(<-ch)
			}
		}
	}()

	// Пишем в канал каждые 50мс, пока не пришёл сигнал о завершение
	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("timeOut")
			return
		default:
			ch <- i
		}
		time.Sleep(time.Millisecond * 50)
	}
}
