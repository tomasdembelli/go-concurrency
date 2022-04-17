package character_counter

import (
	"sync"
)

func count(i string, what string) int {
	var result int
	for _, s := range i {
		if string(s) == what {
			result++
		}
	}
	return result
}

type counter struct {
	sync.Mutex
	cnt int
}

func (c *counter) add(amount int) {
	c.Lock()
	defer c.Unlock()
	c.cnt += amount
}

func (c *counter) get() int {
	c.Lock()
	defer c.Unlock()
	return c.cnt
}

func viaSharedMem(ii []string, what string) int {
	var c counter
	var wg sync.WaitGroup
	for _, i := range ii {
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			c.add(count(i, what))
		}(i)
	}
	wg.Wait()
	return c.get()
}

func viaChan(ii []string, what string) int {
	ch := make(chan int, len(ii))
	var wg sync.WaitGroup
	for _, i := range ii {
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			ch <- count(i, what)
		}(i)
	}
	wg.Wait()
	var total int
	for {
		select {
		case c := <-ch:
			total += c
		default:
			return total
		}
	}
}
