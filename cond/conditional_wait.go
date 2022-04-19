package cond

import (
	"fmt"
	"sync"
	"time"
)

// UseCond is the replica of the example in Page 54.
func UseCond() {
	tasks := [10]int{1,2,3,4,5,6,7,8,9,10}
	processQueue := make([]int,0,2)

	cl := sync.NewCond(&sync.Mutex{})

	removeQ := func(delay time.Duration) {
		time.Sleep(delay)
		cl.L.Lock()
		fmt.Printf("removing from queue: %d\n", processQueue[0])
		processQueue = processQueue[1:]
		cl.L.Unlock()
		cl.Signal()
	}

	for _, t := range tasks {
		cl.L.Lock()
		for len(processQueue) == 2 {
			cl.Wait()
		}
		fmt.Printf("adding to queue: %d\n", t)
		processQueue = append(processQueue, t)
		go removeQ(time.Millisecond)
		cl.L.Unlock()
	}
}