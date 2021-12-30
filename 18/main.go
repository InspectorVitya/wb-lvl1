/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) get() int64 {
	return atomic.LoadInt64(&c.count)
}

func NewCounter() *Counter {
	return &Counter{}
}

func main() {
	wg := sync.WaitGroup{}
	counter := NewCounter()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.inc()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Result: %d\n", counter.get())
}
