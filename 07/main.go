/*
	Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
)

type concurrentMap struct {
	sync.RWMutex
	data map[int]int
}

func (c *concurrentMap) write(key, val int) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = val
}
func (c *concurrentMap) read(key int) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	res, ok := c.data[key]
	return res, ok
}

func main() {
	mp := concurrentMap{
		RWMutex: sync.RWMutex{},
		data:    make(map[int]int),
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			mp.write(i, i)
		}()
	}

	wg.Wait()
	fmt.Println(mp.data)
}
