package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()//相当于减1
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var c Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(&c, &wg)
	}

	wg.Wait()
	fmt.Println(c.Count())
}
